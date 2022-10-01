package dto

type Contact []struct {
	Addresses Address `json:"addresses,omitempty"`
	Birthday  string  `json:"birthday,omitempty"`
	Emails    Email   `json:"emails,omitempty"`
	Name      Name    `json:"name,omitempty"`
	Org       Org     `json:"org,omitempty"`
	Phones    Phone   `json:"phones,omitempty"`
	Urls      Url     `json:"urls,omitempty"`
}

type Address []struct {
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Type        string `json:"type,omitempty"`
}

type Email []struct {
	Email string `json:"email,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Name struct {
	FormattedName string `json:"formatted_name,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	Suffix        string `json:"suffix,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
}

type Org struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type Phone []struct {
	Phone string `json:"phone,omitempty"`
	Type  string `json:"type,omitempty"`
	WaId  string `json:"wa_id,omitempty"`
}

type Url []struct {
	Url  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}
