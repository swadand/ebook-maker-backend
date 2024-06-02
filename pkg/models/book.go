package models

type Book struct {
	ID         string   `json:"id"`
	Filename   string   `json:"filename"`
	OwnerID    string   `json:"owner_id"`
	SharedToID []string `json:"shared_to"`
}
