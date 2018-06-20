package notify


type Notify struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	UIDs    []string `json:"uids"`
}