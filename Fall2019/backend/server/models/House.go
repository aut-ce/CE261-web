package models

type House struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
	Area     int     `json:"area"`
	Bedrooms int     `json:"bedrooms"`
	Parkings int     `json:"parkings"`
	Location *struct {
		Locality string  `json:"locality"`
		Lat      float64 `json:"lat"`
		Long     float64 `json:"long"`
	} `json:"location"`
	Breadcrumb *[]struct {
		Name string `json:"name"`
	} `json:"breadcrumb"`
	Pic *struct {
		Number int      `json:"number"`
		Image  string   `json:"image"`
		Images []string `json:"images"`
	} `json:"pic"`
	Estate *struct {
		Name  string `json:"name"`
		Logo  string `json:"logo"`
		Phone string `json:"phone"`
	} `json:"estate"`
	Star       bool `json:"star"`
	Bookmarked bool `json:"bookmarked"`
	CreatedAt  int  `json:"created_at"`
}
