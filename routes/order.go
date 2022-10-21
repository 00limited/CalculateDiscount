package routes

import (
	"counting_discount/handlers"
	"counting_discount/package/mysql"
	"counting_discount/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	userRepository := repositories.RepositoryUsers(mysql.DB)
	h := handlers.HandlerOrder(orderRepository, userRepository)

	r.HandleFunc("/users", h.FindOrder).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetOrder).Methods("GET")
	r.HandleFunc("/user", h.CreateOrder).Methods("POST")
	// r.HandleFunc("/user/{id}", h.).Methods("PATCH")
	// r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
