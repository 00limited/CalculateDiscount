package models

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" form:"name" gorm:"type: varchar(100)"`
	Total     int       `json:"total" form:"total" gorm:"type: int"`
	Product   []Product `json:"product" gorm:"many2many:product_user"`
	ProductID []int     `json:"-" form:"product_id" gorm:"-"`
	OrderID   uint      `json:"order_id" gorm:"type: int"`
	Order     Order     `json:"order" gorm:"HasOne"`
}
