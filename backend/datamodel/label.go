package datamodel

/*
Label struct
*/
type Label struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Color        string `json:"color,omitempty"`
	Scope        string `json:"scope"`
	ProjectID    int    `json:"project_id"`
	CreationTime string `json:"creation_time,omitempty"`
	UpdateTime   string `json:"update_time,omitempty"`
	Deleted      bool   `json:"deleted,omitempty"`
}
