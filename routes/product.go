package routes

import (
	"counting_discount/handlers"
	"counting_discount/package/mysql"
	"counting_discount/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	ProductRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(ProductRepository)

	r.HandleFunc("/products", h.FindProduct).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/product", h.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", h.UpdateProduct).Methods("PATCH")
	r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
}
