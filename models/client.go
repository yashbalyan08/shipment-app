package models

type Client struct {
	ClientId   string `json:"clientid,omitempty"`
	ClientName string `json:"clientname,omitempty"`
}
