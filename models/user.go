package models

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Porduct []Product `json:"Order"`
}
