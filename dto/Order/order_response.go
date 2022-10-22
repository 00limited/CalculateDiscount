package orderdto

import "counting_discount/models"

type OrderResponse struct {
	ID          int           `json:"id" gorm:"primary_key:auto_increment"`
	Discount    int           `json:"discount" form:"discount" gorm:"type: int"`
	Total       int           `json:"total" form:"total" gorm:"type: int"`
	MaxDiscount int           `json:"maxdiscount" form:"maxdiscount"`
	Order       []models.User `json:"order"`
	OrderID     []int         `json:"-" form:"order_id" gorm:"-"`
}
