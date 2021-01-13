package datamodel

/*
ChartResult struct
*/
type ChartResult struct {
	Name               string        `json:"name"`
	AllDockerImages    []DockerImage `json:"allDockerImages"`
	Charts             []Chart       `json:"charts"`
	Project            string        `json:"project"`
	LatestVersion      string        `json:"latest_version"`
	Icon               string        `json:"icon"`
	ProjectLab         string        `json:"projectLab,omitempty"`
	OtherV             []DockerImage `json:"otherv"`
	LatestDockerImages []DockerImage `json:"latestDockerImages"`
}

/*
ArtifactWithCharts struct
*/
type ArtifactWithCharts struct {
	Name               string   `json:"name"`
	AllDockerImages    []string `json:"allDockerImages"`
	Charts             []Chart  `json:"charts"`
	Project            string   `json:"project"`
	LatestVersion      string   `json:"latest_version"`
	Icon               string   `json:"icon"`
	ProjectLab         string   `json:"projectLab,omitempty"`
	OtherV             []string `json:"otherv"`
	LatestDockerImages []string `json:"latestDockerImages"`
}

/*
ArtifactWithDockerImages struct
*/
type ArtifactWithDockerImages struct {
	Name               string        `json:"name"`
	AllDockerImages    []DockerImage `json:"allDockerImages"`
	Charts             []string      `json:"charts"`
	Project            string        `json:"project"`
	LatestVersion      string        `json:"latest_version"`
	Icon               string        `json:"icon"`
	ProjectLab         string        `json:"projectLab,omitempty"`
	OtherV             []DockerImage `json:"otherv"`
	LatestDockerImages []DockerImage `json:"latestDockerImages"`
}

/*
Artifact struct
*/
type Artifact struct {
	Name               string   `json:"name"`
	AllDockerImages    []string `json:"allDockerImages"`
	Charts             []string `json:"charts"`
	Project            string   `json:"project"`
	LatestVersion      string   `json:"latest_version"`
	Icon               string   `json:"icon"`
	ProjectLab         string   `json:"projectLab,omitempty"`
	OtherV             []string `json:"otherv"`
	LatestDockerImages []string `json:"latestDockerImages"`
}

/*
StarredArtifacts struct
*/
type StarredArtifacts struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
