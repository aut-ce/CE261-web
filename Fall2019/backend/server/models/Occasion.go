package models

type Occasion struct {
	Section string `json:"section,omitempty"`
	Action  struct {
		Text string `json:"text,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"action,omitempty"`
	Items []House `json:"items,omitempty"`
}
