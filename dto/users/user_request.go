package usersdto

type CreateUserRequest struct {
	Name      string `json:"name" form:"name" `
	ProductId []int  `gorm:"type: int" json:"product" `
}

type UpdateUserRequest struct {
	Name      string `json:"name" form:"name"`
	ProductId []int  `gorm:"type: int" json:"product" `
}
