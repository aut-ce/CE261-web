package models

type Occasion struct {
	Section string `json:"section"`
	Action  *struct {
		Text string `json:"text"`
		URL  string `json:"url"`
	} `json:"action,omitempty"`
	Items []OccasionHouse `json:"items"`
}

type OccasionHouse struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
	Area     int     `json:"area"`
	Bedrooms int     `json:"bedrooms"`
	Location *struct {
		Locality string `json:"locality"`
	} `json:"location"`
	Pic *struct {
		Number int    `json:"number"`
		Image  string `json:"image"`
	} `json:"pic"`
	Estate *struct {
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"estate"`
	Star       bool `json:"star"`
	Bookmarked bool `json:"bookmarked"`
	CreatedAt  int  `json:"created_at"`
	Ca         int  `json:"ca,omitempty"`
}
