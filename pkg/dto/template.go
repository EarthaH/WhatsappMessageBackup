package dto

type Template struct {
	Name       string    `json:"name,omitempty"`
	Language   Language  `json:"language,omitempty"`
	Components Component `json:"component,omitempty"`
	Namespace  string    `json:"namespace,omitempty"`
}

type Language struct {
	Policy string `json:"policy,omitempty"`
	Code   string `json:"code,omitempty"`
}

type Component []struct {
	Type       string             `json:"type,omitempty"`
	SubType    string             `json:"sub_type,omitempty"`
	Parameters ComponentParameter `json:"parameters,omitempty"`
	Index      int                `json:"index,omitempty"`
}

type ComponentParameter []struct {
	Type     string `json:"type,omitempty"`
	Text     string `json:"text,omitempty"`
	Currency string `json:"currency,omitempty"`
	DateTime string `json:"date_time,omitempty"`
	Image    Media  `json:"image,omitempty"`
	Document Media  `json:"document,omitempty"`
	Video    Media  `json:"video,omitempty"`
}

type Button struct {
	SubType    string          `json:"sub_type,omitempty"`
	Index      string          `json:"index,omitempty"`
	Parameters ButtonParameter `json:"parameters,omitempty"`
}

type ButtonParameter []struct {
	Type    string `json:"type,omitempty"`
	Payload string `json:"payload,omitempty"`
	Text    string `json:"text,omitempty"`
}
