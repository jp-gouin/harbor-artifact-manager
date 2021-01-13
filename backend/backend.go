package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"backend/args"
	"backend/database"

	"backend/datamodel"
	dmh "backend/datamodelhandlers"
	"backend/deliveries"
	"backend/notification"

	"backend/helpers"

	"backend/hub"

	"github.com/alexflint/go-arg"
	"github.com/bndr/gojenkins"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jtblin/go-ldap-client"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/grpclog"
)

/*
ChartInfoEntry struct
*/
type ChartInfoEntry struct {
	Name          string `json:"name"`
	TotalVersions int    `json:"total_versions"`
	LatestVersion string `json:"latest_version"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	Icon          string `json:"icon"`
	Home          string `json:"home"`
	Deprecated    bool   `json:"deprecated"`
	Project       string `json:"project"`
}

/*
Config struct for config label (approve label)
*/
type Config struct {
	ConfigLabel datamodel.Label   `json:"configlabel"`
	Projects    []datamodel.Label `json:"projects"`
	ID          string            `json:"id"`
}

type RegAuth struct {
	URL      string `json:"url"`
	Password string `json:"password"`
	User     string `json:"user"`
	Project  string `json:"project"`
}

/*
HarborArtifact struct
*/
type HarborArtifact struct {
	Digest       string                 `json:"digest"`
	ScanOverview map[string]interface{} `json:"scan_overview"`
	Labels       []datamodel.Label      `json:"labels"`
	Created      string                 `json:"push_time"`
	Tags         []struct {
		Name   string `json:"name"`
		Signed bool   `json:"signed"`
	} `json:"tags"`
}

var configLabel datamodel.Label

var wshub *hub.Hub

var dao *database.DAO

var nm *notification.Manager

var ldapCli *ldap.LDAPClient

func getChartList(w http.ResponseWriter, r *http.Request) {
	filter := mux.Vars(r)["filter"]
	quick := mux.Vars(r)["quick"]
	nm.Log(fmt.Sprintf("Get chart list filter %s quick %s", filter, quick))
	projects := dmh.GetAllRepositories()
	chartsinfo := processGetChartList(projects, filter)
	if quick != "true" {
		result := make([]datamodel.ArtifactWithCharts, 0)
		resultFull := make([]datamodel.ChartResult, 0)
		for _, ci := range chartsinfo {
			if filter == "" {
				arti := dmh.GetArtifactWithChart(ci.Name, dao)
				if arti != nil {
					result = append(result, *arti)
				}
			} else {
				arti := dmh.GetArtifact(ci.Name, dao)
				if arti != nil {
					resultFull = append(resultFull, *arti)
				}
			}
		}
		if len(result) != 0 {
			json.NewEncoder(w).Encode(result)
		} else {
			json.NewEncoder(w).Encode(resultFull)
		}

	} else {
		result := make([]datamodel.Artifact, 0)
		for _, ci := range chartsinfo {
			arti := dmh.GetArtifactDB(ci.Name, dao)
			if arti != nil {
				result = append(result, *arti)
			}
		}
		json.NewEncoder(w).Encode(result)
	}
}

/*Get ADS_Validated label from Harbor*/
func getConfig(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", args.Args.Harbor+"/api/v2.0/labels?name=ADS_Validated&scope=g", nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		nm.Log(fmt.Sprintf("there was an error performing the http request +%+v", err))
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var labels []datamodel.Label

	json.Unmarshal([]byte(body), &labels)
	configLabel = labels[0]

	req, err = http.NewRequest("GET", args.Args.Harbor+"/api/v2.0/labels?name=Version_&scope=g", nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		nm.Log(fmt.Sprintf("there was an error performing the http request +%+v", err))
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &labels)
	// TODO add the ID in the input arg
	var config = Config{configLabel, labels, "1001-0000-1001-aigh"}

	json.NewEncoder(w).Encode(config)
}

func postChartData(w http.ResponseWriter, r *http.Request) {
	var chartResult datamodel.ChartResult
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &chartResult)
	dmh.HandlePostLabel(chartResult, dao)
	w.WriteHeader(http.StatusOK)
}

func containLabel(labels []datamodel.Label) bool {
	for _, l := range labels {
		if l.Name == configLabel.Name {
			return true
		}
	}
	return false
}

func getChartData(w http.ResponseWriter, r *http.Request) {
	var chartInfo ChartInfoEntry
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &chartInfo)
	json.NewEncoder(w).Encode(processGetChartData(chartInfo))
}
func updateProjectMembers(w http.ResponseWriter, r *http.Request) {
	var project *datamodel.ProjectDBFull
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &project)
	user, err := helpers.GetUsername(r)
	if err == nil {
		err = dmh.HandleUpdateProjectMembers(project, user, dao, nm)
	}
	if err != nil {
		if err.Error() == "Not Allowed" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func notificateValidation(w http.ResponseWriter, r *http.Request) {
	nm.Log("\n Notify all users for validation \n")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var msg string
	json.Unmarshal(reqBody, &msg)
	user, _ := helpers.GetUsername(r)
	nm.BroadcastMessage(msg, user)
}

func GetAllLabelProject(w http.ResponseWriter, r *http.Request) {
	//TODO get all project against the database instead of harbor -> but need to change the ui and need to get ProjectFull

	user, err := helpers.GetUsername(r)
	projects := dmh.HandleGetProjects(user, dao)

	req, err := http.NewRequest("GET", args.Args.Harbor+"/api/v2.0/labels?name=Project_&scope=g", nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		nm.Log(fmt.Sprintf("there was an error performing the http request +%+v \n", err))
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var labels []datamodel.Label
	//result := make([]GuiProject, 0)
	json.Unmarshal([]byte(body), &labels)
	for p, project := range projects {
		for v, version := range project.Versions {
			for _, label := range labels {
				// Project_Backends_Version_V1
				if version.Name == label.Name {
					projects[p].Versions[v].Label = label
				}
			}

		}
	}
	json.NewEncoder(w).Encode(&projects)
}
func GetProjectDetails(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var versionStr string
	json.Unmarshal(reqBody, &versionStr)
	nm.Log(fmt.Sprintf("input label %s", versionStr))
	version, err := dmh.GetVersion(versionStr, dao)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No data found in ETCD for project : " + versionStr))
		return
	} else {
		nm.Log(fmt.Sprintf("data %+v ", version))
		json.NewEncoder(w).Encode(version)
	}
}
func RemoveProject(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var label datamodel.Label
	json.Unmarshal(reqBody, &label)
	nm.Log(fmt.Sprintf("input label %+v", label))
	user, err := helpers.GetUsername(r)
	if err == nil {
		err = dmh.HandleRemoveProject(label, user, dao, nm)
	}
	if err != nil {
		if err.Error() == "Not Allowed" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
func CreateNewProject(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var project string
	json.Unmarshal(reqBody, &project)
	user, err := helpers.GetUsername(r)
	if err == nil {
		err = dmh.HandleCreateNewProject(project, user, dao, nm)
	}
	if err != nil {
		if err.Error() == "Not Allowed" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		} else {
			nm.Log(fmt.Sprintf("there was an error performing the http request %+v \n", err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/*
AddProjectToArtifact : Add a project to an artifact
*/
func AddProjectToArtifact(w http.ResponseWriter, r *http.Request) {
	var chartResult datamodel.ChartResult
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &chartResult)
	//user, err := helpers.GetUsername(r)
	/*if err == nil {
		err = dmh.HandleAddProjectToArtifact(ctx, chartResult, user, kv, nm)
	}
	if err != nil {
		if err.Error() == "Not Allowed" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}*/
}

/*
AddProjectsToArtifact : Add a project to an artifact
*/
func AddProjectsToArtifact(w http.ResponseWriter, r *http.Request) {
	var versionProject *datamodel.VersionDB
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &versionProject)
	user, err := helpers.GetUsername(r)
	if err == nil {
		err = dmh.HandleAddProjectToArtifact(versionProject, user, dao, nm)
		if err != nil {
			if err.Error() == "Not Allowed" {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			/*for _, chartResult := range chartsResult {
			err = dmh.HandleAddProjectToArtifact(ctx, chartResult, user, kv, nm)
			if err != nil {
				if err.Error() == "Not Allowed" {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(err.Error()))
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
				}
			}*/
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
	}

	if err == nil {
		w.WriteHeader(http.StatusOK)
	}
}

/*
RemoveProjectsToArtifact : Add a project to an artifact
*/
func RemoveProjectsToArtifact(w http.ResponseWriter, r *http.Request) {
	var chartsResult []datamodel.ChartResult
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &chartsResult)
	user, err := helpers.GetUsername(r)
	if err == nil {
		for _, chartResult := range chartsResult {
			err = dmh.HandleRemoveProjectToArtifact(chartResult, user, dao, nm)
			if err != nil {
				if err.Error() == "Not Allowed" {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(err.Error()))
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
	}

	if err == nil {
		w.WriteHeader(http.StatusOK)
	}
}

/*
RemoveProjectToArtifact : Remove a project from an artifact
*/
func RemoveProjectToArtifact(w http.ResponseWriter, r *http.Request) {
	var chartResult datamodel.ChartResult
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &chartResult)
	user, err := helpers.GetUsername(r)
	if err == nil {
		err = dmh.HandleRemoveProjectToArtifact(chartResult, user, dao, nm)
	}
	if err != nil {
		if err.Error() == "Not Allowed" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetUserScope(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["user"]
	if helpers.IsAdmin(user) {
		json.NewEncoder(w).Encode("1001-0000-1001-aigh")
	} else {
		json.NewEncoder(w).Encode("")
	}
}

func refreshDownloadLink(w http.ResponseWriter, r *http.Request) {
	project := mux.Vars(r)["project"]
	url := helpers.GetSignedURL(project + ".tar.gz")
	en := json.NewEncoder(w)
	en.SetEscapeHTML(false)
	en.Encode(url)
}
func getAnnoucement(w http.ResponseWriter, r *http.Request) {
	user, _ := helpers.GetUsername(r)
	notifications, err := dao.GetNotifications("Annoucement", 1, user)
	if err == nil {
		json.NewEncoder(w).Encode(notifications)
	}
}
func createAnnoucement(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message string
	json.Unmarshal(reqBody, &message)
	nm.CreateAnnoncement(message)
}
func starArtifact(w http.ResponseWriter, r *http.Request) {
	user, _ := helpers.GetUsername(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var artifact string
	json.Unmarshal(reqBody, &artifact)
	dmh.StarArtifact(artifact, user, dao)
}
func getStarredArtifacts(w http.ResponseWriter, r *http.Request) {
	starred, err := dao.GetStarredArtifacts()
	if err == nil {
		json.NewEncoder(w).Encode(starred)
	}
}
func getUserInfo(w http.ResponseWriter, r *http.Request) {
	user, _ := helpers.GetUsername(r)
	userInfo, err := dmh.GetOrCreateAUser(user, dao, false)
	if err == nil {
		json.NewEncoder(w).Encode(userInfo)
	}
}
func getNotifications(w http.ResponseWriter, r *http.Request) {
	user, _ := helpers.GetUsername(r)
	notifications, err := dao.GetNotifications("Notification", 100, user)
	if err == nil {
		json.NewEncoder(w).Encode(notifications)
	}
}

/*
TestETCD test the database
*/
func TestETCD(w http.ResponseWriter, r *http.Request) {
	val := dao.TestDAO()
	var result interface{}
	json.Unmarshal(val, result)
	nm.Log(fmt.Sprintf("Test etcd : %v", result))
}

func replicateLabels(w http.ResponseWriter, r *http.Request) {
	nm.Log("start label replication")
	decoder := json.NewDecoder(r.Body)
	var reg RegAuth
	err := decoder.Decode(&reg)
	if err == nil {
		dmh.ReplicateLabels(dao.Ctx, reg.URL, reg.User, reg.Password, reg.Project, dao.Kv)
	} else {
		nm.Log(fmt.Sprintf("error %+v", err))
	}

}

func main() {
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	arg.MustParse(args.Args)
	router := mux.NewRouter()

	ctx := context.TODO()
	dao = database.Init(ctx)

	if args.Args.LdapEnable {
		ldapCli = initLdapClient()
	}
	// Init the Hub
	wshub = hub.NewHub()
	go wshub.Run()
	var jenkins *gojenkins.Jenkins
	if args.Args.JenkinsURL != "" {
		// Init Jenkins
		jenkins = gojenkins.CreateJenkins(nil, args.Args.JenkinsURL, args.Args.JenkinsUser, args.Args.JenkinsPass)
	}

	// Init the Notification Manager
	nm = notification.Init(dao, wshub)
	go nm.StartHealthCheck(jenkins, ldapCli)

	// Init the delivery
	deliver := new(deliveries.Deliveries)
	deliver.Init(dao, wshub, nm, jenkins)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "SignIn"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// It's important that this is before your catch-all route ("/")

	// Secure API for Vue app
	sapi := router.PathPrefix("/sapi/v1/").Subrouter()
	sapi.Handle("/getChartList", AuthMiddleware(http.HandlerFunc(getChartList))).Queries("filter", "{filter:.*?}", "quick", "{quick}").Methods("GET")
	sapi.Handle("/getConfig", AuthMiddleware(http.HandlerFunc(getConfig))).Methods("GET")
	sapi.Handle("/getProjects", AuthMiddleware(http.HandlerFunc(GetAllLabelProject))).Methods("GET")
	sapi.Handle("/getUserScope", AuthMiddleware(http.HandlerFunc(GetUserScope))).Queries("user", "{user}").Methods("GET")
	sapi.HandleFunc("/reloadMainDB", ReloadMainDB).Queries("filter", "{filter}").Methods("GET")
	sapi.Handle("/getChartData", AuthMiddleware(http.HandlerFunc(getChartData))).Methods("POST")
	sapi.Handle("/postChartData", AuthMiddleware(http.HandlerFunc(postChartData))).Methods("POST")
	sapi.Handle("/postProject", AuthMiddleware(http.HandlerFunc(CreateNewProject))).Methods("POST")
	sapi.Handle("/generateDelivery", AuthMiddleware(http.HandlerFunc(deliver.GenerateDeliveries))).Methods("POST")
	sapi.Handle("/addProjectToArtifact", AuthMiddleware(http.HandlerFunc(AddProjectToArtifact))).Methods("POST")
	sapi.Handle("/addProjectsToArtifact", AuthMiddleware(http.HandlerFunc(AddProjectsToArtifact))).Methods("POST")
	sapi.Handle("/removeProjectToArtifact", AuthMiddleware(http.HandlerFunc(RemoveProjectToArtifact))).Methods("POST")
	sapi.Handle("/removeProjectsToArtifact", AuthMiddleware(http.HandlerFunc(RemoveProjectsToArtifact))).Methods("POST")
	sapi.Handle("/getProjectDetail", AuthMiddleware(http.HandlerFunc(GetProjectDetails))).Methods("POST")
	sapi.Handle("/removeProject", AuthMiddleware(http.HandlerFunc(RemoveProject))).Methods("POST")
	sapi.Handle("/updateProjectMembers", AuthMiddleware(http.HandlerFunc(updateProjectMembers))).Methods("POST")
	sapi.Handle("/notificateValidation", AuthMiddleware(http.HandlerFunc(notificateValidation))).Methods("POST")
	sapi.Handle("/refreshDownloadLink", AuthMiddleware(http.HandlerFunc(refreshDownloadLink))).Queries("project", "{project}").Methods("GET")
	sapi.Handle("/getAnnoucement", AuthMiddleware(http.HandlerFunc(getAnnoucement))).Methods("GET")
	sapi.Handle("/createAnnoucement", AuthMiddleware(http.HandlerFunc(createAnnoucement))).Methods("POST")
	sapi.Handle("/getNotifications", AuthMiddleware(http.HandlerFunc(getNotifications))).Methods("GET")
	sapi.Handle("/starArtifact", AuthMiddleware(http.HandlerFunc(starArtifact))).Methods("POST")
	sapi.Handle("/getStarredArtifacts", AuthMiddleware(http.HandlerFunc(getStarredArtifacts))).Methods("GET")
	sapi.Handle("/getUserInfo", AuthMiddleware(http.HandlerFunc(getUserInfo))).Methods("GET")
	sapi.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(wshub, w, r)
	})
	// Not secure api for React app
	api := router.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/getChartList", getChartList).Methods("GET")
	api.HandleFunc("/getConfig", getConfig).Methods("GET")
	api.HandleFunc("/getProjects", GetAllLabelProject).Methods("GET")
	api.HandleFunc("/reloadMainDB", ReloadMainDB).Queries("filter", "{filter}").Methods("GET")
	api.HandleFunc("/getChartData", getChartData).Methods("POST")
	api.HandleFunc("/postChartData", postChartData).Methods("POST")
	api.HandleFunc("/postProject", CreateNewProject).Methods("POST")
	api.HandleFunc("/generateDelivery", deliver.GenerateDeliveries).Methods("POST")
	api.HandleFunc("/addProjectToArtifact", AddProjectToArtifact).Methods("POST")
	api.HandleFunc("/addProjectsToArtifact", AddProjectsToArtifact).Methods("POST")
	api.HandleFunc("/removeProjectToArtifact", RemoveProjectToArtifact).Methods("POST")
	api.HandleFunc("/getProjectDetail", GetProjectDetails).Methods("POST")
	api.HandleFunc("/removeProject", RemoveProject).Methods("POST")

	api.HandleFunc("/replicateLabels", replicateLabels).Methods("POST")

	api.HandleFunc("/token", TokenHandler).Methods("POST")
	apitest := router.PathPrefix("/api/test/").Subrouter()
	apitest.HandleFunc("/etcd", TestETCD).Methods("GET")
	// Optional: Use a custom 404 handler for our API paths.
	// api.NotFoundHandler = JSONNotFound
	router.HandleFunc("/token", TokenHandler)
	http.ListenAndServe(":"+strconv.Itoa(args.Args.Listen), handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
