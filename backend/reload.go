package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"backend/args"
	"backend/datamodel"
	dmh "backend/datamodelhandlers"
	"backend/helpers"

	"github.com/gorilla/mux"
)

func processGetChartData(chartInfo ChartInfoEntry) datamodel.ChartResult {
	var result datamodel.ChartResult
	result.Name = chartInfo.Name
	result.Icon = chartInfo.Icon
	result.LatestVersion = chartInfo.LatestVersion
	result.Project = chartInfo.Project
	result.Charts = getAllTag(chartInfo)
	result.AllDockerImages, result.OtherV, result.LatestDockerImages = getAllDockerImages(result.Charts, result.Project)
	return result
}
func getAllTag(chartInfo ChartInfoEntry) []datamodel.Chart {
	req, err := http.NewRequest("GET", args.Args.Harbor+"/api/chartrepo/"+chartInfo.Project+"/charts/"+chartInfo.Name, nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		nm.Log(fmt.Sprintf("there was an error performing the http request %+v ", err))
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var chartsVersion []map[string]interface{}
	json.Unmarshal([]byte(body), &chartsVersion)
	var result = make([]datamodel.Chart, 0)
	//var wg sync.WaitGroup
	//wg.Add(len(chartsVersion))
	for _, c := range chartsVersion {
		//go func(c map[string]interface{}) {
		//defer wg.Done()
		chartDetail := getChartDetail(chartInfo.Project, chartInfo.Name, c["version"].(string))
		chartDetail.CurrentDockerImages = setCurrentDockerImage(chartDetail)
		chartDetail.Values = nil
		result = append(result, chartDetail)
		//}(cv)
	}
	//wg.Wait()
	return result
}
func setCurrentDockerImage(chartDetail datamodel.Chart) []datamodel.DockerImage {
	for key, value := range chartDetail.Values {
		if strings.Contains(key, "repository") {
			if value == nil {
				continue
			}
			var tag = chartDetail.Values[strings.Replace(key, "repository", "tag", 1)]
			var repo string = value.(string)
			if chartDetail.Values["global.imageRegistry"] != nil && chartDetail.Values["global.imageRegistry"] != "" {
				repo = strings.SplitAfterN(chartDetail.Values["global.imageRegistry"].(string), "/", 2)[1] + "/" + repo
				// TODO check if really necessary to split
			}
			// If the repo is hosted in harbor , then it start with the project and since harbor 2.0 the
			// project is no more part of the repo url (e.g i3s/backends/psql is in the i3s project and the repo is backends/psql)
			project := ""
			if strings.Contains(repo, "/") {
				project = strings.SplitN(repo, "/", 2)[0]
				repo = strings.SplitAfterN(repo, "/", 2)[1]
			}
			di := datamodel.DockerImage{chartDetail.Name, nil, nil, repo, tag.(string), "", "", project}
			exist := false
			for _, dock := range chartDetail.CurrentDockerImages {
				if dock.Repository == repo && dock.Tag == tag.(string) {
					exist = true
					break
				}
			}
			if !exist {
				chartDetail.CurrentDockerImages = append(chartDetail.CurrentDockerImages, di)
			}
		}
	}
	return chartDetail.CurrentDockerImages
}
func getChartDetail(project string, name string, tag string) datamodel.Chart {
	req, err := http.NewRequest("GET", args.Args.Harbor+"/api/chartrepo/"+project+"/charts/"+name+"/"+tag, nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		nm.Log(fmt.Sprintf("there was an error performing the http request +%+v ", err))
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var chartDetail datamodel.Chart
	json.Unmarshal([]byte(body), &chartDetail)
	nm.Log(fmt.Sprintf("lookup %s", args.Args.Harbor+"/api/chartrepo/"+project+"/charts/"+name+"/"+tag))
	if chartDetail.Metadata["version"] == nil {
		nm.Log(fmt.Sprintf(" Error on chart %s / %s / %s  retry request", project, name, tag))
		nm.Log(fmt.Sprintf(" Chartdetail : %+v", chartDetail))
		return getChartDetail(project, name, tag)
	}
	chartDetail.Version = chartDetail.Metadata["version"].(string)
	chartDetail.Name = chartDetail.Metadata["name"].(string)
	chartDetail.Project = project
	return chartDetail

}
func getAllDockerImages(charts []datamodel.Chart, project string) ([]datamodel.DockerImage, []datamodel.DockerImage, []datamodel.DockerImage) {
	if len(charts) == 0 {
		return nil, nil, nil
	}
	var result = make([]datamodel.DockerImage, 0)
	var otherv = make([]datamodel.DockerImage, 0)
	var latest = make([]datamodel.DockerImage, 0)
	var imgs []string
	for _, chart := range charts {
		//var wg sync.WaitGroup
		//wg.Add(len(chart.CurrentDockerImages))
		for i, d := range chart.CurrentDockerImages {
			//go func(d DockerImage) {
			//	defer wg.Done()
			if !helpers.Contains(imgs, d.Repository) {
				var page = 0
				for {
					imgs = append(imgs, d.Repository)
					nm.Log(fmt.Sprintf("request  %s", args.Args.Harbor+"/api/v2.0/projects/"+project+"/repositories/"+d.Repository+"/artifacts"))
					req, err := http.NewRequest("GET", args.Args.Harbor+"/api/v2.0/projects/"+project+"/repositories/"+strings.ReplaceAll(d.Repository, "/", "%252F")+"/artifacts?with_tag=true&with_scan_overview=true&with_label=true&page_size=100&page="+strconv.Itoa(page), nil)
					req.SetBasicAuth(args.Args.User, args.Args.Password)
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						nm.Log(fmt.Sprintf("there was an error performing the http request %+v ", err))
					}
					defer resp.Body.Close()
					body, _ := ioutil.ReadAll(resp.Body)
					var harborArtifacts []HarborArtifact
					json.Unmarshal([]byte(body), &harborArtifacts)
					if len(harborArtifacts) == 0 {
						break
					}
					var dockerImages []datamodel.DockerImage
					for _, ha := range harborArtifacts {
						for _, tag := range ha.Tags {
							di := datamodel.DockerImage{tag.Name, ha.Labels, ha.ScanOverview, d.Repository, tag.Name, ha.Created, ha.Digest, project}
							dockerImages = append(dockerImages, di)
							if d.Tag == di.Tag {
								chart.CurrentDockerImages[i].Digest = di.Digest
							}
						}
					}
					for i, dock := range dockerImages {
						dockerImages[i].Repository = d.Repository
						dockerImages[i].Tag = dock.Name
						result = append(result, dockerImages[i])
						isCandidate, errCandid := isCandidateForOtherV(charts, dockerImages[i])
						if isCandidate && errCandid == nil {
							otherv = append(otherv, dockerImages[i])
						}
					}
					ldi, ldierr := helpers.GetLatest(dockerImages)
					if ldierr == nil {
						latest = append(latest, ldi...)
						chart.LatestDockerImages = append(latest, ldi...)
					}
					page++
				}
			}
			//}(di)
		}
		//wg.Wait()
	}
	return result, otherv, latest
}

func processGetChartList(projects []datamodel.Repositories, filter string) []ChartInfoEntry {
	var result = make([]ChartInfoEntry, 0)
	for _, p := range projects {
		req, err := http.NewRequest("GET", args.Args.Harbor+"/api/chartrepo/"+p.Name+"/charts", nil)
		req.SetBasicAuth(args.Args.User, args.Args.Password)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			nm.Log(fmt.Sprintf("there was an error performing the http request +%+v ", err))
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var charts []ChartInfoEntry
		json.Unmarshal([]byte(body), &charts)
		for _, c := range charts {
			c.Project = p.Name
			if strings.Compare(filter, "") != 0 && strings.Compare(filter, c.Name) == 0 {
				result = append(result, c)
			} else if strings.Compare(filter, "") == 0 {
				result = append(result, c)
			}
		}
	}
	return result
}

func processReloadMainDB(filter string) {
	nm.Log("Loading main database from harbor data")
	nm.BroadcastMessage("Database is releoading, do not use the app until done", "bot")

	nm.Log("etcd is fine")
	admins := strings.Split(args.Args.Admins, ",")
	repos := dmh.GetAllRepositories()
	fmt.Printf("repositories : +%v \n", repos)
	chartsInfo := processGetChartList(repos, filter)
	chartR := make([]datamodel.ChartResult, 0)

	//	var wg sync.WaitGroup
	//	wg.Add(len(chartsInfo))
	for _, ci := range chartsInfo {
		nm.Log("iterate on " + ci.Name)
		//		go func(chartinfo ChartInfoEntry) {
		//			defer wg.Done()
		cr := processGetChartData(ci)
		chartR = append(chartR, cr)
		dmh.SetArtifact(cr, true, dao)
		//		}(ci)
	}
	//	wg.Wait()
	for _, myC := range chartR {
		for _, c := range myC.Charts {
			for _, l := range c.Labels {
				if strings.Contains(l.Name, "Project_") && !strings.Contains(l.Name, "#") {
					version, _ := dao.GetVersionDB(l.Name)
					// if the project does not exist on etcd then create it and put the data
					if version == nil {
						nm.Log("do no exist in etcd")
						// Create new project if it does not exist and the label version using func in project
						_, version, _ := dmh.CreateNewProject(l.Name, admins[0], dao)
						version.Charts = append(version.Charts, dmh.GetChartID(c))
						b, err := json.Marshal(version)
						if err != nil {
							nm.Log(err.Error())
							break
						}
						dao.Set(l.Name, string(b))
					} else {
						nm.Log("exist in etcd")
						version.Charts = appendChartIfNotExist(version.Charts, c)
						b, err := json.Marshal(version)
						if err != nil {
							nm.Log(err.Error())
							break
						}
						dao.Set(l.Name, string(b))
					}
				}
			}
		}
		for _, dock := range myC.AllDockerImages {
			for _, l := range dock.Labels {
				if strings.Contains(l.Name, "Project_") && !strings.Contains(l.Name, "#") {
					version, _ := dao.GetVersionDB(l.Name)
					// if the project does not exist on etcd then create it and put the data
					if version == nil {
						nm.Log("do no exist in etcd \n Won't do anything because the dockerimage should be attached to a chart")
					} else {
						nm.Log("exist in etcd")
						// Check if the DI is already contained in a chart in this project
						version.DockerImages = appendDockerImageIfNotExistInProject(version.DockerImages, dock, version.Charts)
						b, err := json.Marshal(version)
						if err != nil {
							nm.Log(err.Error())
							break
						}
						dao.Set(l.Name, string(b))
					}
				}
			}
		}
	}
	nm.BroadcastMessage("Reload done ! ", "bot")
}
func appendChartIfNotExist(charts []string, chart datamodel.Chart) []string {
	found := false
	cID := dmh.GetChartID(chart)
	for _, c := range charts {
		if c == cID {
			found = true
			break
		}
	}
	nm.Log(fmt.Sprintf("%v", found))
	if found {
		return charts
	}
	result := append(charts, cID)
	return result
}
func appendDockerImageIfNotExistInProject(dockerImages []string, dock datamodel.DockerImage, charts []string) []string {

	for _, cID := range charts {
		c := dmh.GetChart(cID, dao)
		for _, do := range c.CurrentDockerImages {
			if dmh.GetDockerImageID(do) == dmh.GetDockerImageID(dock) {
				return dockerImages
			}
		}
	}
	return appendDockerImageIfNotExist(dockerImages, dock)
}
func appendDockerImageIfNotExist(dockerImages []string, dock datamodel.DockerImage) []string {
	diID := dmh.GetDockerImageID(dock)
	for _, di := range dockerImages {
		if di == diID {
			nm.Log("Docker image already exist")
			return dockerImages
		}
	}
	result := append(dockerImages, diID)
	return result
}

/*
ReloadMainDB : Reload the database by querying through all data
*/
func ReloadMainDB(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["filter"]
	nm.Log("filter : " + id)
	processReloadMainDB(id)
	w.WriteHeader(http.StatusOK)
}
