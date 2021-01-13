package datamodelhandlers

import (
	"fmt"
	"strings"

	"backend/args"
	dm "backend/datamodel"
)

func isUserAllowedOnProject(project *dm.ProjectDB, username string, ownerAction bool) bool {
	fmt.Printf("Check if user %s is authorized to modify project %+v", username, project)
	if isAdmin(username) {
		return true
	}
	for _, owner := range project.Owners {
		if owner == username {
			return true
		}
	}
	if ownerAction {
		return false
	}
	for _, member := range project.Members {
		if member == username {
			return true
		}
	}
	return false
}

func isAdmin(username string) bool {
	admins := strings.Split(args.Args.Admins, ",")
	for _, admin := range admins {
		if admin == username {
			return true
		}
	}
	return false
}
