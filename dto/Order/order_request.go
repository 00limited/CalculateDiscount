package orderdto

type OrderRequest struct {
	Discount    int   `json:"discount"`
	Maxdiscount int   `json:"maxdiscount"`
	UserId      []int `json:"userid"`
}
