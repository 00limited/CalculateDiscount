package productdto

type CreateProductrRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Price  int    `json:"price" form:"price" validate:"required"`
	Quanty int    `json:"quanty" form:"quanty" validate:"required"`
}

type UpdateProductRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Price  int    `json:"price" form:"price" validate:"required"`
	Quanty int    `json:"quanty" form:"quanty" validate:"required"`
}
