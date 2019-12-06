package models

type Magazine struct {
	Section string `json:"section,omitempty"`
	Items   []struct {
		ID    string `json:"id,omitempty"`
		Image string `json:"image,omitempty"`
		Title string `json:"title,omitempty"`
	} `json:"items,omitempty"`
}
