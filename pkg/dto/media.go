package dto

type Media struct {
	Id       string   `json:"id,omitempty"`
	Link     string   `json:"link,omitempty"`
	Caption  string   `json:"caption,omitempty"`
	Filename string   `json:"filename,omitempty"`
	Provider Provider `json:"provider,omitempty"`
}

type Provider struct {
	Name string `json:"name,omitempty"`
}
