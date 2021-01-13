package datamodelhandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"backend/database"
	notif "backend/notification"

	dm "backend/datamodel"

	"backend/args"
)

/*
GetAllRepositories from Harbor
*/
func GetAllRepositories() []dm.Repositories {
	req, err := http.NewRequest("GET", args.Args.Harbor+"/api/v2.0/projects", nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("there was an error performing the http request +%+v \n", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var repos []dm.Repositories
	json.Unmarshal([]byte(body), &repos)
	return repos
}

// HandleAddProjectToArtifact add project to an artifact
func HandleAddProjectToArtifact(version *dm.VersionDB, username string, dao *database.DAO, nm *notif.Manager) error {
	fmt.Printf("project %+v \n", version)

	projectName := strings.SplitN(version.Name, "_", 3)[1]
	project, _ := dao.GetProjectDB(projectName)
	allowed := isUserAllowedOnProject(project, username, true)
	if !allowed {
		nm.NotifyProject("Add Artifact to "+projectName+" forbidden ", username, "Error", "Artifact", project)
		return errors.New("Not Allowed")
	}
	previousVersion, _ := dao.GetVersionDB(version.Name)
	HandlePostLabelFromVersion(version, previousVersion, dao)
	b, err := json.Marshal(version)
	if err != nil {
		return err
	}
	dao.Set(version.Name, string(b))
	nm.NotifyProject("Add Artifact to "+projectName, username, "Success", "Artifact", project)
	return nil
}

/*
AddProjectToEtcdDB : add data to ETCD
*/
func AddProjectToEtcdDB(project dm.ProjectDB, chartResult dm.ChartResult, dao *database.DAO) error {
	version, _ := dao.GetVersionDB(chartResult.ProjectLab)
	if version == nil {
		version = &dm.VersionDB{chartResult.ProjectLab, make([]string, 0), make([]string, 0), "", dm.Label{}, ""}
		project.Versions = append(project.Versions, version.Name)
		b, err := json.Marshal(project)
		if err != nil {
			fmt.Println(err)
			return err
		}
		dao.Set("TProject_"+project.Name, string(b))
	}
	if len(chartResult.AllDockerImages) > 0 {
		for _, di := range chartResult.AllDockerImages {
			// Found example
			_, found := find(version.DockerImages, GetDockerImageID(di))
			if !found {
				version.DockerImages = append(version.DockerImages, GetDockerImageID(di))
			}
		}
	}
	if len(chartResult.Charts) > 0 {
		for _, c := range chartResult.Charts {
			// Found example
			_, found := find(version.Charts, GetChartID(c))
			if !found {
				version.Charts = append(version.Charts, GetChartID(c))
			}
		}
	}
	version.LastLink = ""
	b, err := json.Marshal(version)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dao.Set(chartResult.ProjectLab, string(b))
	return nil
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		fmt.Printf("\n %s vs %s \n", item, val)
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func GetVersion(version string, dao *database.DAO) (*dm.Version, error) {
	versionDB, _ := dao.GetVersionDB(version)
	if versionDB == nil {
		return nil, errors.New("No Version found")
	}
	versionResult := dm.Version{version, make([]dm.Chart, 0), make([]dm.DockerImage, 0), versionDB.LastLink, versionDB.Status}

	for _, diID := range versionDB.DockerImages {
		di, _ := dao.GetDockerImage(diID)
		versionResult.DockerImages = append(versionResult.DockerImages, *di)
	}
	for _, cID := range versionDB.Charts {
		c := GetChart(cID, dao)
		versionResult.Charts = append(versionResult.Charts, c)
	}
	return &versionResult, nil
}

func HandleRemoveProjectToArtifact(chartResult dm.ChartResult, username string, dao *database.DAO, nm *notif.Manager) error {
	fmt.Println("\n Remove Project \n")
	version, _ := dao.GetVersionDB(chartResult.ProjectLab)

	fmt.Printf("\n init charts %+v \n", version.Charts)
	fmt.Printf("\n init dis %+v \n", version.DockerImages)

	projectName := strings.SplitN(chartResult.ProjectLab, "_", 3)[1]
	project, _ := dao.GetProjectDB("TProject_" + projectName)

	allowed := isUserAllowedOnProject(project, username, true)
	if !allowed {
		nm.NotifyProject("Remove artifact of "+projectName+" forbidden", username, "Error", "Artifact", project)
		return errors.New("Not Allowed")
	}
	HandlePostLabel(chartResult, dao)
	if len(chartResult.AllDockerImages) > 0 {
		for _, di := range chartResult.AllDockerImages {
			i, found := find(version.DockerImages, GetDockerImageID(di))
			if found {
				version.DockerImages = remove(version.DockerImages, i)
			}
		}
		nm.Log(fmt.Sprintf("\n new dockerimages %+v \n", version.DockerImages))
	}
	if len(chartResult.Charts) > 0 {
		for _, c := range chartResult.Charts {
			i, found := find(version.Charts, GetChartID(c))
			if found {
				version.Charts = remove(version.Charts, i)
			}
		}
		nm.Log(fmt.Sprintf("\n new charts %s \n", version.Charts))
	}
	version.LastLink = ""
	b, err := json.Marshal(version)
	if err == nil {
		dao.Set(version.Name, string(b))
	}
	nm.NotifyProject("Remove artifact of "+projectName, username, "Success", "Artifact", project)
	return err
}
func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
func HandleRemoveProject(label dm.Label, username string, dao *database.DAO, nm *notif.Manager) error {
	nm.Log("Remove Project : " + label.Name)
	projectName := strings.SplitN(label.Name, "_", 3)[1]
	nm.Log(projectName)
	project, _ := dao.GetProjectDB("TProject_" + projectName)

	var err error
	allowed := isUserAllowedOnProject(project, username, true)
	if !allowed {
		nm.NotifyProject("Deletion of "+projectName+" forbidden ", username, "Error", "Project Deletion", project)
		return errors.New("Not Allowed")
	}
	DeleteLabel(args.Args.Harbor + "/api/v2.0/labels/" + strconv.Itoa(label.ID))

	dao.Delete(label.Name)

	newList := make([]string, 0)
	for _, version := range project.Versions {
		nm.Log(version + "  vs   " + label.Name)
		if version != label.Name {
			newList = append(newList, version)
		}
	}
	if len(newList) == 0 {
		dao.Delete("TProject_" + projectName)
		nm.NotifyProject("Deletion of "+projectName, username, "Success", "Project Deletion", project)
		return err
	}
	nm.Log(fmt.Sprintf("new list : %+v", newList))
	project.Versions = newList
	b, err := json.Marshal(project)
	if err == nil {
		dao.Set("TProject_"+projectName, string(b))
	}
	nm.NotifyProject("Deletion of "+projectName, username, "Success", "Project Deletion", project)
	return err
}

/*
HandleCreateNewProject create a project and a first version corresponding to the input data
Create all entry in ETCD and create a label in harbor
*/
func HandleCreateNewProject(project string, username string, dao *database.DAO, nm *notif.Manager) error {
	label := dm.Label{0, project, "Automated creation of : " + project, "", "g", 0, "", "", false}
	AddLabel(args.Args.Harbor+"/api/v2.0/labels", label)
	projectDB, _, err := CreateNewProject(project, username, dao)
	if err == nil {
		nm.NotifyProject("Creation of "+project, username, "Success", "Project Creation", projectDB)
	} else {
		nm.NotifyProject("Fail creation of "+project, username, "Error", "Fail Project Creation", projectDB)
	}
	return err
}

/*
CreateNewProject create a project and a first version corresponding to the input data
Create all entry in ETCD
*/
func CreateNewProject(version string, username string, dao *database.DAO) (*dm.ProjectDB, *dm.VersionDB, error) {

	projectName := strings.SplitN(version, "_", 3)[1]
	projectDB, _ := dao.GetProjectDB("TProject_" + projectName)
	fmt.Printf("creation of project : TProject_%s", projectName)
	fmt.Printf("\n %v ", projectDB)
	if projectDB == nil {
		projectDB = &dm.ProjectDB{projectName, make([]string, 0), make([]string, 0), make([]string, 0)}
		projectDB.Owners = append(projectDB.Owners, username)
		fmt.Printf("creation of project : %+v", projectDB)
	}
	versionDB, _ := dao.GetVersionDB(version)
	if versionDB != nil {
		return nil, nil, errors.New("Already Exist")
	}
	versionDB = &dm.VersionDB{version, make([]string, 0), make([]string, 0), "", dm.Label{}, ""}
	projectDB.Versions = append(projectDB.Versions, version)

	b, err := json.Marshal(projectDB)
	dao.Set("TProject_"+projectName, string(b))

	b, err = json.Marshal(versionDB)
	dao.Set(version, string(b))

	return projectDB, versionDB, err
}

/*
HandleGetProjects get project and resolve version
*/
func HandleGetProjects(username string, dao *database.DAO) []*dm.ProjectDBFull {

	projectsDB, _ := dao.GetProjectsDB("TProject_")

	var result []*dm.ProjectDBFull
	for _, projectDB := range projectsDB {
		isMember := isUserAllowedOnProject(projectDB, username, false)
		if isMember {
			projectDBFull := dm.ProjectDBFull{projectDB.Name, make([]dm.VersionDB, 0), projectDB.Owners, projectDB.Members}
			for _, versionStr := range projectDB.Versions {
				versionDB, _ := dao.GetVersionDB(versionStr)
				projectDBFull.Versions = append(projectDBFull.Versions, *versionDB)
			}
			sort.SliceStable(projectDBFull.Versions, func(i, j int) bool {
				a, _ := strconv.Atoi(strings.SplitN(projectDBFull.Versions[i].Name, "_", 4)[3])
				b, _ := strconv.Atoi(strings.SplitN(projectDBFull.Versions[j].Name, "_", 4)[3])
				return a < b
			})
			result = append(result, &projectDBFull)
		}
	}
	return result
}

/*
HandleUpdateProjectMembers update project members
*/
func HandleUpdateProjectMembers(project *dm.ProjectDBFull, username string, dao *database.DAO, nm *notif.Manager) error {

	projectDB, _ := dao.GetProjectDB("TProject_" + project.Name)
	allowed := isUserAllowedOnProject(projectDB, username, true)
	if !allowed {
		nm.NotifyProject("Change members of "+project.Name+" forbidden ", username, "Error", "Project Change", projectDB)
		return errors.New("Not Allowed")
	}
	projectDB.Members = project.Members
	projectDB.Owners = project.Owners

	b, err := json.Marshal(projectDB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dao.Set("TProject_"+project.Name, string(b))
	nm.NotifyProject("Change members of "+project.Name, username, "Success", "Project Change", projectDB)
	return nil
}
