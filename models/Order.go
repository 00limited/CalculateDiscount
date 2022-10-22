package models

type Order struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Discount    int    `json:"discount" form:"discount" gorm:"type: int"`
	Total       int    `json:"total" form:"total" gorm:"type: int"`
	MaxDiscount int    `json:"maxdiscount" form:"maxdiscount"`
	Users       []User `json:"users" gorm:"HasMany"`
}
