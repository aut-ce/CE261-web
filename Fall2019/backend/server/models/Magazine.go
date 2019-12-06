package models

type Magazine struct {
	ID    string `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	Title string `json:"title,omitempty"`
}

type MagazineResponse struct {
	Section string     `json:"section,omitempty"`
	Items   []Magazine `json:"items,omitempty"`
}
