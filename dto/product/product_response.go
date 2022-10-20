package productdto

type ProductResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name" form:"name" validate:"required"`
	Price  int    `json:"price" form:"price" validate:"required"`
	Quanty int    `json:"quanty" form:"quanty" validate:"required"`
}
