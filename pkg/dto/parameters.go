package dto

type Parameter []struct {
	Audio            Media       `json:"audio,omitempty"`
	Contacts         Contact     `json:"contacts,omitempty"`
	Context          Context     `json:"context,omitempty"`
	Document         Media       `json:"document,omitempty"`
	Image            Media       `json:"image,omitempty"`
	Interactive      Interactive `json:"interactive,omitempty"`
	Location         Location    `json:"location,omitempty"`
	MessagingProduct string      `json:"messaging_product,omitempty"`
	PreviewUrl       bool        `json:"preview_url,omitempty"`
	RecipientType    string      `json:"recipient_type,omitempty"`
	Status           string      `json:"status,omitempty"`
	Sticker          Media       `json:"sticker,omitempty"`
	Template         Template    `json:"template,omitempty"`
	Text             Text        `json:"text,omitempty"`
	To               string      `json:"to,omitempty"`
	Type             string      `json:"type,omitempty"`
}

type Context struct {
	MessageId string `json:"message_id,omitempty"`
}
