package usersdto

type CreateUserRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	ProductId []int  `gorm:"type: int" json:"productId" validate:"required"`
}

type UpdateUserRequest struct {
	Name      string `json:"name" form:"name"`
	ProductId []int  `gorm:"type: int" json:"productId" validate:"required"`
}
