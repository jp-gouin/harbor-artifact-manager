package datamodelhandlers

import (
	"encoding/json"
	"fmt"

	"backend/database"
	dm "backend/datamodel"
)

/*
SetDockerImage into the database forceUpdate if necessary (when updating the label of the dockerimage)
*/
func SetDockerImage(image dm.DockerImage, forceUpdate bool, dao *database.DAO) string {
	id := GetDockerImageID(image)
	oldimg, _ := dao.GetDockerImage(id)
	if oldimg == nil || forceUpdate {
		b, err := json.Marshal(image)
		if err == nil {
			dao.Set(id, string(b))
		} else {
			fmt.Println(err)
		}
	}
	return id
}

/*
GetDockerImageID return the database id of a DockerImage
*/
func GetDockerImageID(image dm.DockerImage) string {
	return image.Repository + image.Tag
}
