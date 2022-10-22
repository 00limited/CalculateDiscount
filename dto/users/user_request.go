package usersdto

type CreateUserRequest struct {
	Name    string `json:"name" form:"name" `
	Price   int    `json:"price" `
	OrderID uint   `json:"order_id"`
}

type UpdateUserRequest struct {
	Name      string `json:"name" form:"name"`
	ProductId []int  `gorm:"type: int" json:"product" `
	OrderID   uint   `json:"order_id"`
}
