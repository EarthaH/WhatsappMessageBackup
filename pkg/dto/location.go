package dto

type Location struct {
	Longitude string `json:"longitude,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
}
