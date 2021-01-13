package datamodelhandlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"backend/args"
	"backend/database"
	dm "backend/datamodel"

	"go.etcd.io/etcd/clientv3"
)

/*
AddLabel to harbor
*/
func AddLabel(url string, lab dm.Label) {
	addLabel(url, lab, args.Args.User, args.Args.Password)
}

func addLabel(url string, lab dm.Label, harborUser string, harborPass string) {
	fmt.Println("add label to " + url)
	b, err := json.Marshal(lab)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(harborUser, harborPass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("there was an error performing the http request %+v \n", err)
	}
	fmt.Printf("result : %+v", resp)
	defer resp.Body.Close()
}

/*
DeleteLabel on harbor
*/
func DeleteLabel(url string) {
	deleteLabel(url, args.Args.User, args.Args.Password)
}

func deleteLabel(url string, harborUser string, harborPass string) {
	fmt.Println("delete label to " + url)
	req, err := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth(harborUser, harborPass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("there was an error performing the http request %+v \n", err)
	}
	fmt.Printf("response %+v \n", resp)
	defer resp.Body.Close()
}

/*
ReplicateLabels from the current harbor to the other
Assuming that all charts and dockerimages exist on the target Harbor
And only the base url has changed e.g harbor.example.com/... to harbor.other.org/...
Please note that label id will also be replicated, so no labels should be present on target Harbor
*/
func ReplicateLabels(ctx context.Context, urlHarbor2 string, user string, password string, project string, kv clientv3.KV) {
	gr, _ := kv.Get(ctx, "", clientv3.WithPrefix())
	fmt.Printf("get all data from etcd %d", len(gr.Kvs))
	if len(gr.Kvs) != 0 {
		for _, val := range gr.Kvs {
			var v dm.DockerImage
			json.Unmarshal(val.Value, &v)
			if v.Repository != "" {
				for _, label := range v.Labels {
					if !label.Deleted {
						addLabel(urlHarbor2+"/api/v2.0/labels/", label, user, password)
						addLabel(urlHarbor2+"/api/v2.0/repositories/"+v.Repository+"/tags/"+v.Tag+"/labels", label, user, password)
					}
				}
			} else {
				var cdb dm.ChartDB
				json.Unmarshal(val.Value, &cdb)
				if cdb.Version != "" {
					for _, label := range cdb.Labels {
						if !label.Deleted {
							addLabel(urlHarbor2+"/api/v2.0/labels/", label, user, password)
							addLabel(urlHarbor2+"/api/v2.0/chartrepo/"+project+"/charts/"+cdb.Name+"/"+cdb.Version+"/labels", label, user, password)
						}
					}
				}
			}
		}
	}
}

/*
HandlePostLabel handle label add/remove on Harbor
*/
func HandlePostLabel(chartResult dm.ChartResult, dao *database.DAO) {
	for _, chart := range chartResult.Charts {
		for _, label := range chart.Labels {
			if label.Deleted {
				DeleteLabel(args.Args.Harbor + "/api/v2.0/chartrepo/" + chartResult.Project + "/charts/" + chartResult.Name + "/" + chart.Version + "/labels/" + strconv.Itoa(label.ID))
			} else {
				AddLabel(args.Args.Harbor+"/api/v2.0/chartrepo/"+chartResult.Project+"/charts/"+chartResult.Name+"/"+chart.Version+"/labels", label)
			}
		}
		SetChart(chart, true, dao)
	}
	for _, di := range chartResult.AllDockerImages {
		for _, label := range di.Labels {
			if label.Deleted {
				DeleteLabel(args.Args.Harbor + "/api/v2.0/projects/" + chartResult.Project + "/repositories/" + strings.ReplaceAll(di.Repository, "/", "%252F") + "/artifacts/" + di.Digest + "/labels/" + strconv.Itoa(label.ID))
			} else {
				AddLabel(args.Args.Harbor+"/api/v2.0/projects/"+chartResult.Project+"/repositories/"+strings.ReplaceAll(di.Repository, "/", "%252F")+"/artifacts/"+di.Digest+"/labels", label)
			}
		}
		SetDockerImage(di, true, dao)
	}
}

/*
HandlePostLabelFromVersion handle label add/remove on Harbor
*/
func HandlePostLabelFromVersion(version *dm.VersionDB, previousVersion *dm.VersionDB, dao *database.DAO) {
	listNew := getDiffItem(version.Charts, previousVersion.Charts)
	listRemove := getDiffItem(previousVersion.Charts, version.Charts)

	for _, cNewID := range listNew {
		chart := GetChart(cNewID, dao)
		AddLabel(args.Args.Harbor+"/api/v2.0/chartrepo/"+chart.Project+"/charts/"+chart.Name+"/"+chart.Version+"/labels", version.Label)
		// Either i change the label of the item now and update the data in the database
		// Or i do not change it now and it will natively be updated during the reload
		// Since it's the versionDB that is read by the generator it's not vital to update the data now but it's cleaner...
		//chart.Labels = append(chart.Labels, version.Label)
		//SetChart(chart, true, dao)
	}
	for _, cRemID := range listRemove {
		chart := GetChart(cRemID, dao)
		DeleteLabel(args.Args.Harbor + "/api/v2.0/chartrepo/" + chart.Project + "/charts/" + chart.Name + "/" + chart.Version + "/labels/" + strconv.Itoa(version.Label.ID))
		// See previous comment
		//SetChart(chart, true, dao)
	}
	listNew = getDiffItem(version.DockerImages, previousVersion.DockerImages)
	listRemove = getDiffItem(previousVersion.DockerImages, version.DockerImages)
	for _, dNewID := range listNew {
		fmt.Println("\n\n ID :::: " + dNewID)
		di, _ := dao.GetDockerImage(dNewID)
		AddLabel(args.Args.Harbor+"/api/v2.0/projects/"+di.Project+"/repositories/"+strings.ReplaceAll(di.Repository, "/", "%252F")+"/artifacts/"+di.Digest+"/labels", version.Label)
		// See previous comment
		//SetDockerImage(*di, true, dao)
	}
	for _, dRemID := range listRemove {
		di, _ := dao.GetDockerImage(dRemID)
		DeleteLabel(args.Args.Harbor + "/api/v2.0/projects/" + di.Project + "/repositories/" + strings.ReplaceAll(di.Repository, "/", "%252F") + "/artifacts/" + di.Digest + "/labels/" + strconv.Itoa(version.Label.ID))
		// See previous comment
		//SetDockerImage(*di, true, dao)
	}
}
func getDiffItem(listA []string, listB []string) []string {
	result := make([]string, 0)
	for _, elemA := range listA {
		found := false
		for _, elemB := range listB {
			if elemA == elemB {
				found = true
				break
			}
		}
		if !found {
			result = append(result, elemA)
		}
	}
	return result
}
