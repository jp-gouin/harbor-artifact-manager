package datamodel

/*
Project struct
*/
type Repositories struct {
	ID   int                    `json:"project_id"`
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"-"`
}

type GuiProject struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
}

/*
ProjectDB struct
*/
type ProjectDB struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
	Owners   []string `json:"owners"`
	Members  []string `json:"members"`
}

/*
ProjectDBFull struct
*/
type ProjectDBFull struct {
	Name     string      `json:"name"`
	Versions []VersionDB `json:"versions"`
	Owners   []string    `json:"owners"`
	Members  []string    `json:"members"`
}

/*
VersionDB struct
*/
type VersionDB struct {
	Name         string   `json:"name"`
	Charts       []string `json:"charts"`
	DockerImages []string `json:"dockerImages"`
	LastLink     string   `json:"lastLink"`
	Label        Label    `json:"label"` // This field are only fill on getProjects call from client , it should be form when a project is created with the label
	Status       string   `json:"status"`
}

/*
Version struct
*/
type Version struct {
	Name         string        `json:"name"`
	Charts       []Chart       `json:"charts"`
	DockerImages []DockerImage `json:"dockerImages"`
	LastLink     string        `json:"lastLink"`
	Status       string        `json:"status"`
}

/*
ProjectFull struct
*/
type ProjectFull struct {
	Name     string    `json:"name"`
	Versions []Version `json:"versions"`
	Owners   []string  `json:"owners"`
	Members  []string  `json:"members"`
}
