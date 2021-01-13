package helpers

import (
	"backend/args"
	"backend/datamodel"
	"strings"
)

/*
GetbaseURL get the docker uri for the given docker image
*/
func GetbaseURL(di *datamodel.DockerImage) string {
	baseurl := strings.Replace(args.Args.Harbor, "/api/", "/", 1)
	baseurl = strings.Replace(baseurl, "https://", "", 1)
	if di.Project != "" {
		baseurl = baseurl + "/" + di.Project + "/"
	}
	return baseurl + di.Repository + ":" + di.Tag
}
