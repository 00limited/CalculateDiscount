package models

type Product struct {
	ID     int    `json:"id" `
	Name   string `json:"name"`
	Quanty int    `json:"quanty"`
	Price  int    `json:"price"`
}
