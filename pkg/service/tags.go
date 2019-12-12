package service

// Tags represents a diff of tags
type Tags struct {
	OldTags []string `json:"oldTags"`
	NewTags []string `json:"newTags"`
}
