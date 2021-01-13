package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"backend/args"
	"backend/datamodel"
)

func isCandidateForOtherV(charts []datamodel.Chart, dockerImage datamodel.DockerImage) (bool, error) {
	var valid = true
	exit, err := isImageValid(dockerImage.Repository)
	exit2, err2 := isImageValid(dockerImage.Tag)
	if err != nil || *exit || err2 != nil || *exit2 {
		return false, err
	}
	for _, chart := range charts {
		if chart.Metadata["appVersion"] == nil {
			return false, errors.New("No AppVersion")
		} else if strings.Contains(dockerImage.Tag, ".") {
			if strings.HasPrefix(chart.Metadata["appVersion"].(string), strings.Split(dockerImage.Tag, ".")[0]) {
				return false, nil
			}
		} else if strings.Contains(dockerImage.Tag, "-") {
			if strings.HasPrefix(chart.Metadata["appVersion"].(string), strings.Split(dockerImage.Tag, "-")[0]) {
				return false, nil
			}
		} else {
			if strings.Contains(chart.Metadata["appVersion"].(string), dockerImage.Tag) {
				return false, nil
			}
		}
	}
	return valid, nil
}

func isImageValid(value string) (*bool, error) {
	matched, err := regexp.MatchString(args.Args.Regexp, value)

	if err != nil {
		fmt.Printf("error applying regexp reason : %+v \n ", err)
		return nil, err
	}
	return &matched, nil
}
