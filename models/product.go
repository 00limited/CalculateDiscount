package models

type Product struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Name   string `json:"name" form:"name" gorm:"type: varchar(100)"`
	Quanty int    `json:"quanty" form:"qty"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
}
