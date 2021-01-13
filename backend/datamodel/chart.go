package datamodel

/*
Chart struct , chart data returned to the client
*/
type Chart struct {
	Name                string                 `json:"name"`
	Version             string                 `json:"version"`
	Metadata            map[string]interface{} `json:"metadata"`
	Security            map[string]interface{} `json:"security"`
	Labels              []Label                `json:"labels"`
	Values              map[string]interface{} `json:"values"`
	CurrentDockerImages []DockerImage          `json:"currentDockerImages"`
	Urls                []string               `json:"urls"`
	AppVersion          string                 `json:"appVersion"`
	Dependencies        []interface{}          `json:"dependencies"`
	LatestDockerImages  []DockerImage          `json:"latestDockerImages"`
	Project             string                 `json:"project,omitempty"`
}

/*
ChartDB struct , chart data stored in the database
*/
type ChartDB struct {
	Name                string                 `json:"name"`
	Version             string                 `json:"version"`
	Metadata            map[string]interface{} `json:"metadata"`
	Security            map[string]interface{} `json:"security"`
	Labels              []Label                `json:"labels"`
	Values              map[string]interface{} `json:"values"`
	CurrentDockerImages []string               `json:"currentDockerImages"`
	Urls                []string               `json:"urls"`
	AppVersion          string                 `json:"appVersion"`
	Dependencies        []interface{}          `json:"dependencies"`
	LatestDockerImages  []string               `json:"latestDockerImages"`
	Project             string                 `json:"project,omitempty"`
}
