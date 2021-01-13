package datamodel

// Notification struct handled by the hub
type Notification struct {
	ID       string  `json:"id"`
	Date     string  `json:"date"`
	Title    string  `json:"title"`
	Owner    string  `json:"owner"`
	Type     string  `json:"type"`
	Payload  string  `json:"payload"`
	Progress float64 `json:"progress,omitempty"`
	Severity string  `json:"severity"`
	WSType   string  `json:"mutation"`
}
