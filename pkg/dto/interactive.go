package dto

type Interactive struct {
	Type   string `json:"type,omitempty"`
	Header Header `json:"header,omitempty"`
	Body   Body   `json:"body,omitempty"`
	Footer Footer `json:"footer,omitempty"`
	Action Action `json:"action,omitempty"`
}

type Header struct {
	Document Media  `json:"document,omitempty"`
	Image    Media  `json:"image,omitempty"`
	Video    Media  `json:"video,omitempty"`
	Text     string `json:"text,omitempty"`
	Type     string `json:"type,omitempty"`
}

type Body struct {
	Text string `json:"text,omitempty"`
}

type Footer struct {
	Text string `json:"text,omitempty"`
}

type Action struct {
	Button            string       `json:"button,omitempty"`
	Buttons           ActionButton `json:"buttons,omitempty"`
	Sections          Section      `json:"section,omitempty"`
	CatalogId         string       `json:"catalog_id,omitempty"`
	ProductRetailerId string       `json:"product_retailer_id,omitempty"`
}

type ActionButton []struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Id    string `json:"id,omitempty"`
}

type Section []struct {
	Title        string      `json:"title,omitempty"`
	Rows         Row         `json:"rows,omitempty"`
	ProductItems ProductItem `json:"product_items,omitempty"`
}

type Row []struct {
	Title       string `json:"title,omitempty"`
	Id          string `json:"ID,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProductItem []struct {
	ProductRetailerId string `json:"product_retailer_id,omitempty"`
}
