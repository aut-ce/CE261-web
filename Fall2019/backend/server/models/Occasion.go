package models

type Occasion struct {
	Section string `json:"section"`
	Action  *struct {
		Text string `json:"text"`
		URL  string `json:"url"`
	} `json:"action,omitempty"`
	Items []House `json:"items"`
}
