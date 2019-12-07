package models

type Chart struct {
	SelectedColor string       `json:"selected_color,omitempty"`
	OtherColor    string       `json:"other_color,omitempty"`
	Data          []ChartHouse `json:"data"`
}

type ChartHouse struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}
