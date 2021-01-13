package datamodel

/*
DockerImage struct
*/
type DockerImage struct {
	Name         string                 `json:"name"`
	Labels       []Label                `json:"labels"`
	ScanOverview map[string]interface{} `json:"scan_overview"`
	Repository   string                 `json:"repository"`
	Tag          string                 `json:"tag"`
	Created      string                 `json:"created"`
	Digest       string                 `json:"digest"`
	Project      string                 `json:"project"`
}
