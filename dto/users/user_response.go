package usersdto

type UserResponse struct {
	Name      string `json:"name" form:"name" `
	ProductId []int  `gorm:"type: int" json:"productId" `
}
