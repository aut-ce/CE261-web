package models

type Chart struct {
	SelectedColor string  `json:"selected_color,omitempty"`
	OtherColor    string  `json:"other_color,omitempty"`
	Data          []House `json:"data,omitempty"`
}
