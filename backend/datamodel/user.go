package datamodel

// User is a defenition of user
type User struct {
	Username string   `json:"members"`
	Robot    *RobotV1 `json:"robot"`
	Starred  []string `json:"starred"`
}

// Robot holds the details of a robot.
type Robot struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Secret      string        `json:"secret"`
	Salt        string        `json:"-"`
	Duration    int64         `json:"duration"`
	ProjectID   int64         `json:"project_id"`
	ExpiresAt   int64         `json:"expires_at"`
	Disabled    bool          `json:"disabled"`
	Visible     bool          `json:"-"`
	Permissions []*Permission `json:"permissions"`
}

// Permission ...
type Permission struct {
	Kind      string    `json:"kind"`
	Namespace string    `json:"namespace"`
	Access    []*Policy `json:"access"`
	Scope     string    `json:"-"`
}

// Policy the type of policy
type Policy struct {
	Resource string `json:"resource,omitempty"`
	Action   string `json:"action,omitempty"`
	Effect   string `json:"effect,omitempty"`
}

/*
 RobotV1 struct
*/
type RobotV1 struct {
	ID          *int64               `json:"id,omitempty"`
	Name        string               `json:"name"`
	Fullname    string               `json:"fullname"`
	Description string               `json:"description"`
	Token       string               `json:"token,omitempty"`
	ExpiresAt   int64                `json:"expires_at"`
	Access      []RobotAccountAccess `json:"access"`
}

/*
RobotAccountAccess struct
*/
type RobotAccountAccess struct {
	Action   string `json:"action"`
	Resource string `json:"Resource"`
}
