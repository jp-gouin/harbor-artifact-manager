package datamodelhandlers

import (
	"encoding/json"
	"fmt"

	"backend/database"
	dm "backend/datamodel"
)

/*
SetChart into the database, transform the data to ChartDB and set dockerImage in the database
return the id of the element
*/
func SetChart(chart dm.Chart, forceUpdate bool, dao *database.DAO) string {
	id := GetChartID(chart)
	chartDB, _ := dao.GetChartDB(id)
	if chartDB == nil || forceUpdate {
		chartDB := dm.ChartDB{chart.Name, chart.Version, chart.Metadata, chart.Security, chart.Labels, nil, []string{}, chart.Urls, chart.AppVersion, chart.Dependencies, []string{}, chart.Project}
		for _, di := range chart.CurrentDockerImages {
			diID := SetDockerImage(di, forceUpdate, dao)
			chartDB.CurrentDockerImages = append(chartDB.CurrentDockerImages, diID)
		}
		for _, di := range chart.LatestDockerImages {
			diID := SetDockerImage(di, forceUpdate, dao)
			chartDB.LatestDockerImages = append(chartDB.LatestDockerImages, diID)
		}
		b, err := json.Marshal(chartDB)
		if err == nil {
			dao.Set(id, string(b))
		} else {
			fmt.Println(err)
		}
	}
	return id
}

/*
GetChart from the database, fetch a ChartDB and turn it into Chart
*/
func GetChart(id string, dao *database.DAO) dm.Chart {
	chartDB, _ := dao.GetChartDB(id)
	chart := dm.Chart{chartDB.Name, chartDB.Version, chartDB.Metadata, chartDB.Security, chartDB.Labels, nil, make([]dm.DockerImage, 0), chartDB.Urls, chartDB.AppVersion, chartDB.Dependencies, make([]dm.DockerImage, 0), chartDB.Project}

	for _, currentDockerImages := range chartDB.CurrentDockerImages {
		di, _ := dao.GetDockerImage(currentDockerImages)
		chart.CurrentDockerImages = append(chart.CurrentDockerImages, *di)
	}
	for _, latestDockerImages := range chartDB.LatestDockerImages {
		di, _ := dao.GetDockerImage(latestDockerImages)
		chart.CurrentDockerImages = append(chart.CurrentDockerImages, *di)
	}

	return chart
}

/*
GetChartID return the database id of a chart
*/
func GetChartID(chart dm.Chart) string {
	return chart.Name + chart.Version
}
