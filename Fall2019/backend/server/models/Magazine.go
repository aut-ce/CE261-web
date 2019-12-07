package models

type Magazine struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Title string `json:"title"`
}

type MagazineResponse struct {
	Section string     `json:"section"`
	Items   []Magazine `json:"items"`
}
