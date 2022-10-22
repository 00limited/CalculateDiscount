package routes

import (
	"counting_discount/handlers"
	"counting_discount/package/mysql"
	"counting_discount/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUsers(mysql.DB)
	productRepository := repositories.RepositoryProduct(mysql.DB)
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerUser(userRepository, productRepository, orderRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")

}
