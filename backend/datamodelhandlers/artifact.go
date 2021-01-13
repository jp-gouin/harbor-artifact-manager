package datamodelhandlers

import (
	"encoding/json"
	"fmt"

	"backend/database"
	dm "backend/datamodel"
)

/*
SetArtifact into the database, transform the data to ArtifactDB and set dockerImag and chart in the database
*/
func SetArtifact(chart dm.ChartResult, forceUpdate bool, dao *database.DAO) {
	id := chart.Name
	oldArtifact, _ := dao.GetArtifactDB(id)
	if oldArtifact == nil || forceUpdate {
		fmt.Printf("Set artifact %s \n", chart.Name)
		artifactDB := dm.Artifact{chart.Name, []string{}, []string{}, chart.Project, chart.LatestVersion, chart.Icon, chart.ProjectLab, []string{}, []string{}}
		for _, di := range chart.AllDockerImages {
			diID := SetDockerImage(di, forceUpdate, dao)
			artifactDB.AllDockerImages = append(artifactDB.AllDockerImages, diID)
		}
		for _, di := range chart.OtherV {
			diID := SetDockerImage(di, forceUpdate, dao)
			artifactDB.OtherV = append(artifactDB.OtherV, diID)
		}
		for _, di := range chart.LatestDockerImages {
			diID := SetDockerImage(di, forceUpdate, dao)
			artifactDB.LatestDockerImages = append(artifactDB.LatestDockerImages, diID)
		}
		for _, c := range chart.Charts {
			cID := SetChart(c, false, dao)
			artifactDB.Charts = append(artifactDB.Charts, cID)
		}
		b, err := json.Marshal(artifactDB)
		if err == nil {
			dao.Set(id, string(b))
		} else {
			fmt.Println(err)
		}
	}
}

/*
GetArtifact from the database, fetch a ArtifactDB and turn it into ChartResult
*/
func GetArtifact(id string, dao *database.DAO) *dm.ChartResult {
	artifactDB, _ := dao.GetArtifactDB(id)
	chartResult := dm.ChartResult{artifactDB.Name, make([]dm.DockerImage, 0), make([]dm.Chart, 0), artifactDB.Project, artifactDB.LatestVersion, artifactDB.Icon, artifactDB.ProjectLab, make([]dm.DockerImage, 0), make([]dm.DockerImage, 0)}

	for _, id := range artifactDB.AllDockerImages {
		di, _ := dao.GetDockerImage(id)
		chartResult.AllDockerImages = append(chartResult.AllDockerImages, *di)
	}
	for _, id := range artifactDB.OtherV {
		di, _ := dao.GetDockerImage(id)
		chartResult.OtherV = append(chartResult.OtherV, *di)
	}
	for _, id := range artifactDB.LatestDockerImages {
		di, _ := dao.GetDockerImage(id)
		chartResult.LatestDockerImages = append(chartResult.LatestDockerImages, *di)
	}
	for _, id := range artifactDB.Charts {
		c := GetChart(id, dao)
		chartResult.Charts = append(chartResult.Charts, c)
	}

	return &chartResult
}

/*
GetArtifactDB get the artifact but without processing the docker image and full chart data
Used for the dashboard UI to quickly display the page
*/
func GetArtifactDB(id string, dao *database.DAO) *dm.Artifact {
	artifactDB, _ := dao.GetArtifactDB(id)
	return artifactDB
}

/*
GetArtifactWithChart from the database, fetch a ArtifactDB and turn it into ChartResult
param : scope => if all resolve all data , dockerimage resolve only dockerimage , chart resolve only charts
*/
func GetArtifactWithChart(id string, dao *database.DAO) *dm.ArtifactWithCharts {
	artifactDB, _ := dao.GetArtifactDB(id)
	if artifactDB == nil {
		return nil
	}
	artifact := dm.ArtifactWithCharts{artifactDB.Name, artifactDB.AllDockerImages, make([]dm.Chart, 0), artifactDB.Project, artifactDB.LatestVersion, artifactDB.Icon, artifactDB.ProjectLab, artifactDB.OtherV, artifactDB.LatestDockerImages}
	for _, id := range artifactDB.Charts {
		c := GetChart(id, dao)
		artifact.Charts = append(artifact.Charts, c)
	}
	return &artifact
}

/*
	StarArtifact add a star to an Artifact, unstar if already present
*/
func StarArtifact(artifact string, username string, dao *database.DAO) error {
	// Check is the user can star the project
	user, _ := dao.GetUser(username)
	for index, star := range user.Starred {
		// If already star then unstar
		if star == artifact {
			user.Starred = remove(user.Starred, index)
			SaveUser(user, dao)
			allStarred, _ := dao.GetStarredArtifacts()
			artSt := getStared(allStarred, artifact)
			artSt.Count = artSt.Count - 1
			b, _ := json.Marshal(allStarred)
			dao.Set("StarredArtifacts", string(b))
			return nil
		}
	}
	user.Starred = append(user.Starred, artifact)
	SaveUser(user, dao)
	allStarred, err := dao.GetStarredArtifacts()
	if err != nil {
		// Init artifact since none exist
		allStarred = make([]*dm.StarredArtifacts, 0)
		newStar := &dm.StarredArtifacts{artifact, 0}
		allStarred = append(allStarred, newStar)
	}
	star := getStared(allStarred, artifact)
	fmt.Printf("\n\n star : %v", star)
	if star != nil {
		star.Count = star.Count + 1
	} else {
		newStar := &dm.StarredArtifacts{artifact, 1}
		allStarred = append(allStarred, newStar)
	}
	fmt.Printf("\n\n star : %v", star)
	fmt.Printf("\n\n allStarred : %v", allStarred[0])
	b, err := json.Marshal(allStarred)
	dao.Set("StarredArtifacts", string(b))
	return nil
}
func getStared(list []*dm.StarredArtifacts, id string) *dm.StarredArtifacts {
	for i, s := range list {
		if s.Name == id {
			return list[i]
		}
	}
	return nil
}
