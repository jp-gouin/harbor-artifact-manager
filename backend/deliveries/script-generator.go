package deliveries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	b64 "encoding/base64"

	"backend/args"
	"backend/database"
	dm "backend/datamodel"
	dmh "backend/datamodelhandlers"
	"backend/helpers"
	"backend/hub"
	notif "backend/notification"

	"github.com/bndr/gojenkins"
)

// Generalize this to all other component. Init the datamodel with an object and etcd client
// Use this object to call all datamodel function

/*
Deliveries struct to init the package
*/
type Deliveries struct {
	Dao     *database.DAO
	Nm      *notif.Manager
	Jenkins *gojenkins.Jenkins
}

type DataGen struct {
	Name         string   `json:"chartname"`
	URL          string   `json:"charturl"`
	Dockerimages []string `json:"dockerimages"`
}

/*
Init the package
*/
func (d *Deliveries) Init(dao *database.DAO, hub *hub.Hub, nm *notif.Manager, jenkins *gojenkins.Jenkins) {
	d.Dao = dao
	d.Nm = nm
	d.Jenkins = jenkins
}

/*
GenerateDeliveries generate the convinient script to generate offline install
*/
func (d *Deliveries) GenerateDeliveries(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var project map[string]interface{}
	json.Unmarshal(reqBody, &project)
	d.Nm.Log(fmt.Sprintf("input label %s ", project))
	mode := project["mode"].(string)
	force := project["force"].(bool)
	versionName := project["label"].(string)
	version, err1 := dmh.GetVersion(versionName, d.Dao)
	projectName := strings.SplitN(versionName, "_", 3)[1]
	projectDB, err2 := d.Dao.GetProjectDB("TProject_" + projectName)
	username, err3 := helpers.GetUsername(r)
	user, err4 := dmh.GetOrCreateAUser(username, d.Dao, true)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		strerr := err1.Error() + err2.Error() + err3.Error()
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error " + strerr + " with : " + versionName))
		return
	}

	var repoList []string
	var chartList []string
	dataGen := make([]DataGen, 0)

	additionalDG := DataGen{"Additional", "", make([]string, 0)}
	for _, di := range version.DockerImages {
		d.Nm.Log(di.Repository + ":" + di.Tag + "")
		baseurl := helpers.GetbaseURL(&di)
		if !helpers.Contains(additionalDG.Dockerimages, baseurl) {
			additionalDG.Dockerimages = append(additionalDG.Dockerimages, baseurl)
		}
		if !helpers.Contains(repoList, baseurl) {
			repoList = append(repoList, baseurl)
		}
	}
	dataGen = append(dataGen, additionalDG)
	for _, c := range version.Charts {
		url := c.Metadata["urls"].([]interface{})
		dg := DataGen{c.Name, url[0].(string), make([]string, 0)}
		for _, di := range c.CurrentDockerImages {
			baseurl := helpers.GetbaseURL(&di)
			if !helpers.Contains(dg.Dockerimages, baseurl) {
				dg.Dockerimages = append(dg.Dockerimages, baseurl)
			}
			if !helpers.Contains(repoList, baseurl) {
				repoList = append(repoList, baseurl)
			}
		}
		dataGen = append(dataGen, dg)
		if !helpers.Contains(chartList, url[0].(string)) {
			chartList = append(chartList, url[0].(string))
		}
	}

	rlist, _ := json.Marshal(repoList)
	clist, _ := json.Marshal(chartList)
	repoListFormated := strings.Replace(string(rlist), "[", "(", 1)
	repoListFormated = strings.Replace(repoListFormated, "]", ")", 1)
	repoListFormated = strings.Replace(repoListFormated, ",", " ", -1)
	chartListFormated := strings.Replace(string(clist), "[", "(", 1)
	chartListFormated = strings.Replace(chartListFormated, "]", ")", 1)
	chartListFormated = strings.Replace(chartListFormated, ",", " ", -1)

	d.Nm.Log(string(rlist))

	//read, err := ioutil.ReadFile("/Users/jpgouin/Work/Backends/docker/i3s-backends-validator/backend/backends-save-images.sh")
	read, err := ioutil.ReadFile("/usr/share/backends-save-images.sh")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n user %v \n", user)
	newContents := strings.Replace(string(read), "@repoList", repoListFormated, -1)
	newContents = strings.Replace(newContents, "@chartList", chartListFormated, -1)
	newContents = strings.Replace(newContents, "@user", strings.Replace(user.Robot.Fullname, "$", "\\$", -1), -1)
	newContents = strings.Replace(newContents, "@password", user.Robot.Token, -1)
	if args.Args.S3Bucket != "" && args.Args.JenkinsURL != "" && mode == "s3" {
		// Watchdog , if a build is already launch for the version do not execute a new one !
		if version.Status == "pending" && !force {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("A build is already pending : " + versionName))
			return
		}

		// post on Jenkins URL with the script and bucket name and key as parameters
		// use websocket and get the progress of jenkins and log it in the wss
		// Should add the script as file but it's not working see (https://issues.jenkins-ci.org/browse/JENKINS-27413).
		//	param = append(param, jenkinsParameter{"backends-save-images.sh", "", "file0"})
		artefacts, _ := json.Marshal(dataGen)

		var params = map[string]string{
			"bucket":     args.Args.S3Bucket,
			"harboruser": strings.Replace(user.Robot.Fullname, "$", "\\$", -1),
			"harborpass": user.Robot.Token,
			"filename":   versionName + ".tar.gz",
			"script":     b64.StdEncoding.EncodeToString([]byte(newContents)),
			"artefacts":  string(artefacts),
		}
		jobName := strings.SplitN(args.Args.JenkinsBuildAPI, "/", 4)[2]
		job, err := d.Jenkins.GetJob(jobName)
		qid, err := job.InvokeSimple(params)
		if err != nil {
			d.Nm.NotifyProject("Build of "+versionName+" failed", user.Username, "Error", "Build failed", projectDB)
			d.Nm.Log(fmt.Sprintf("Some error triggered while invoking job  %s in queue %d Error %s", job.GetName(), qid, err))
			return
		}
		d.Nm.Log(fmt.Sprintf(" number ::: %d", qid))
		if qid == 0 {
			go d.trackJenkinsBuild(versionName, user.Username, projectDB)
		} else {
			go d.trackJenkinsBuildByID(qid, versionName, user.Username, projectDB)
		}

		//defer resp.Body.Close()
	} else {
		w.Write([]byte(newContents))
	}
}
func (d *Deliveries) trackJenkinsBuildByID(location int64, versionName string, user string, projectDB *dm.ProjectDB) {
	notif, _ := d.Nm.NotifyBuild("Build of "+versionName, user, "Progress", projectDB)
	ticker := time.NewTicker(1 * time.Second)
	task, err := d.Jenkins.GetQueueItem(location)
	go func() {
		for t := range ticker.C {
			_ = t // we don't print the ticker time, so assign this `t` variable to underscore `_` to avoid error
			task.Poll()
			d.Nm.Log(fmt.Sprintf(" tick -> get task from queuid : %d", location))
			if err != nil {
				notif.Type = "project"
				notif.Severity = "Error"
				notif.Title = "Build failed !"
				notif.Payload = err.Error()
				d.Nm.UpdateNotification(notif, projectDB)
				ticker.Stop()
				return
			}
			if task.Raw.Executable.URL != "" {
				// start fct
				go d.getProgress(notif, task.Raw.Executable.Number, task.Raw.Task.Name, versionName, user, projectDB)
				ticker.Stop()
			}
		}
	}()

	// wait for 10 seconds
	time.Sleep(10 * time.Second)
	ticker.Stop()
}
func (d *Deliveries) trackJenkinsBuild(versionName string, user string, projectDB *dm.ProjectDB) {
	notif, _ := d.Nm.NotifyBuild("Build of "+versionName, user, "Progress", projectDB)
	jobName := strings.SplitN(args.Args.JenkinsBuildAPI, "/", 4)[2]
	ticker := time.NewTicker(1 * time.Second)
	// Wait 5 sec to be sure the last build for the job will be ours and not one already running...
	time.Sleep(5 * time.Second)
	go func() {
		for t := range ticker.C {
			_ = t // we don't print the ticker time, so assign this `t` variable to underscore `_` to avoid error
			job, err := d.Jenkins.GetJob(jobName)
			if err != nil {
				notif.Type = "project"
				notif.Severity = "Error"
				notif.Title = "Build failed !"
				notif.Payload = err.Error()
				d.Nm.UpdateNotification(notif, projectDB)
				ticker.Stop()
				return
			}
			b, err := job.GetLastBuild()
			if err != nil {
				notif.Type = "project"
				notif.Severity = "Error"
				notif.Title = "Build failed !"
				notif.Payload = err.Error()
				d.Nm.UpdateNotification(notif, projectDB)
				ticker.Stop()
				return
			}
			currentTime := time.Now()
			bt := b.Raw.Timestamp / 1000

			if b.GetResult() != "SUCCESS" && b.GetResult() != "FAILURE" && b.GetResult() != "ABORTED" && currentTime.Unix()-bt < 10 {
				d.Nm.Log(fmt.Sprintf(" Job is in good state : %+v ", b))
				go d.getProgress(notif, b.GetBuildNumber(), jobName, versionName, user, projectDB)
				ticker.Stop()
			}
		}
	}()

	// wait for 10 seconds
	time.Sleep(20 * time.Second)
	ticker.Stop()
}

func (d *Deliveries) getProgress(notif *dm.Notification, number int64, jobName string, versionName string, user string, projectDB *dm.ProjectDB) {
	ticker := time.NewTicker(5 * time.Second)
	b, err := d.Jenkins.GetBuild(jobName, number)
	versionDB, _ := d.Dao.GetVersionDB(versionName)
	go func() {
		for t := range ticker.C {
			_ = t // we don't print the ticker time, so assign this `t` variable to underscore `_` to avoid error
			b.Poll()
			if err != nil {
				notif.Type = "project"
				notif.Severity = "Error"
				notif.Title = "Build failed !"
				notif.Payload = err.Error()
				versionDB.Status = "Error"
				json, _ := json.Marshal(versionDB)
				d.Dao.Set(versionName, string(json))
				d.Nm.UpdateNotification(notif, projectDB)
				ticker.Stop()
			}
			versionDB.Status = b.Raw.Result
			d.Nm.Log(fmt.Sprintf("Display the build :: %+v", b))
			switch b.Raw.Result {
			case "SUCCESS":
				notif.Type = "project"
				notif.Severity = "Success"
				notif.Title = "Build done !"
				str := fmt.Sprintf("<a href='%s'>Download Link</a>", helpers.GetSignedURL(versionName+".tar.gz"))
				notif.Payload = str
				d.Nm.UpdateNotification(notif, projectDB)
				versionDB.LastLink = str
				json, _ := json.Marshal(versionDB)
				d.Dao.Set(versionName, string(json))
				ticker.Stop()
			case "FAILURE":
				notif.Type = "project"
				notif.Severity = "Error"
				notif.Title = "Build failed"
				d.Nm.UpdateNotification(notif, projectDB)
				ticker.Stop()
			case "ABORTED":
				notif.Type = "project"
				notif.Severity = "Error"
				notif.Title = "Build Aborted"
				d.Nm.UpdateNotification(notif, projectDB)
				ticker.Stop()
			default:
				versionDB.Status = "pending"
				dt := time.Now()
				delta := float64(dt.Unix() - b.Raw.Timestamp/1000)
				p1 := delta / float64(b.Raw.EstimatedDuration)
				//progress := float64((dt.Unix()-(b.Raw.Timestamp/1000))/b.Raw.EstimatedDuration) * 100 * 100 * 10
				d.Nm.Log(fmt.Sprintf(" delta :: %f", delta))
				d.Nm.Log(fmt.Sprintf(" p1 %f", p1))
				notif.Progress = p1 * 100 * 10
				d.Nm.UpdateNotification(notif, projectDB)
			}
			json, _ := json.Marshal(versionDB)
			d.Dao.Set(versionName, string(json))
			d.Nm.UpdateNotification(notif, projectDB)
		}
	}()
}
