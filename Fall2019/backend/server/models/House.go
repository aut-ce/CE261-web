package models

type House struct {
	ID       string  `json:"id"`
	Title    string  `json:"title,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Type     string  `json:"type,omitempty"`
	Area     int     `json:"area,omitempty"`
	Bedrooms int     `json:"bedrooms,omitempty"`
	Parkings int     `json:"parkings,omitempty"`
	Location struct {
		Locality string  `json:"locality,omitempty"`
		Lat      float64 `json:"lat,omitempty"`
		Long     float64 `json:"long,omitempty"`
	} `json:"location,omitempty"`
	Breadcrumb []struct {
		Name string `json:"name,omitempty"`
	} `json:"breadcrumb,omitempty"`
	Pic struct {
		Number int      `json:"number,omitempty"`
		Images []string `json:"images,omitempty"`
	} `json:"pic,omitempty"`
	Estate struct {
		Name  string `json:"name,omitempty"`
		Logo  string `json:"logo,omitempty"`
		Phone string `json:"phone,omitempty"`
	} `json:"estate,omitempty"`
	Star       bool `json:"star,omitempty"`
	Bookmarked bool `json:"bookmarked,omitempty"`
	CreatedAt  int  `json:"created_at,omitempty"`
}
