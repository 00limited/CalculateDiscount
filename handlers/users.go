package handlers

import (
	dto "counting_discount/dto/result"
	usersdto "counting_discount/dto/users"
	"counting_discount/models"
	"counting_discount/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository    repositories.UserRepository
	ProductRepository repositories.ProductRepository
}

func HandlerUser(UserRepository repositories.UserRepository, ProductRepository repositories.ProductRepository) *handlerUser {
	return &handlerUser{UserRepository, ProductRepository}
}

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var user models.User
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseUser(user)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := usersdto.CreateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(request)

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ := h.UserRepository.FindProductById(request.ProductId)

	var totalPrice int
	for i := 1; i <= len(request.ProductId); i++ {
		pddetail, _ := h.ProductRepository.GetProduct(i)
		totalPrice += pddetail.Price
	}

	user := models.User{
		Name:    request.Name,
		Total:   totalPrice,
		Product: product,
	}

	user, err = h.UserRepository.CreateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, _ = h.UserRepository.GetUser(user.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get product id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var productId []int
	for _, r := range r.FormValue("categoryId") {
		if int(r-'0') >= 0 {
			productId = append(productId, int(r-'0'))
		}
	}

	request := usersdto.UpdateUserRequest{
		Name:      r.FormValue("name"),
		ProductId: productId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get all product data by id []
	var product []models.Product
	if len(productId) != 0 {
		product, _ = h.UserRepository.FindProductById(productId)
	}

	user, _ := h.UserRepository.GetUser(id)

	user.Name = request.Name
	user.Product = product

	user, err = h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get product id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	DeleteUser, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: DeleteUser}
	json.NewEncoder(w).Encode(response)
}

func convertResponseUser(p models.User) models.User {
	return models.User{
		ID:      p.ID,
		Name:    p.Name,
		Total:   p.Total,
		Product: p.Product,
	}
}
