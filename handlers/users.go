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
	OrderRepository   repositories.OrderRepository
}

func HandlerUser(UserRepository repositories.UserRepository, ProductRepository repositories.ProductRepository, OrderRepository repositories.OrderRepository) *handlerUser {
	return &handlerUser{UserRepository, ProductRepository, OrderRepository}
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

	requests := []usersdto.CreateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(&requests)
	if err != nil {
		fmt.Println("koko")
		fmt.Println(err)
	}

	// fmt.Println(requests)
	var totalOrder int
	var discountUser float64
	var countDis int
	var finalPrice int

	for i := 0; i < len(requests); i++ {

	}
	// fmt.Println("totalOrder", totalOrder)

	for i := 0; i < len(requests); i++ {
		request := requests[i]

		//counting all price for all buyer
		totalOrder += request.Price

		//counting persent each user from total order
		discountUser = (float64(request.Price) / float64(totalOrder)) * 100

		//counting discount for order and check maxdiscount no more than 30.000
		if (totalOrder*30)/100 > 30000 {
			countDis = 30000
		} else {
			countDis = (totalOrder * 30) / 100
		}

		//counting finalprice after total old order substrack by countdis
		finalPrice = totalOrder - countDis

		//update and counting for each user mush be pay
		UserTotal := (float64(finalPrice) * discountUser) / 100
		requests[i].Price = int(UserTotal)

		// fmt.Println("discountUser", discountUser)
		// fmt.Println("countDis: ", countDis)
		// fmt.Println("finalPrice: ", finalPrice)
		// fmt.Println(UserTotal)
		// fmt.Printf("user %s, final price yg harus dibayar %d\n", request.Name, request.Price)
	}

	orderStruct := models.Order{
		Discount:    30,
		Total:       totalOrder,
		MaxDiscount: 30000,
	}

	// create order di sini

	order, err := h.OrderRepository.CreateOrder(orderStruct)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i := 0; i < len(requests); i++ {
		request := requests[i]

		user := models.User{
			Name:    request.Name,
			Total:   request.Price,
			OrderID: uint(order.ID),
		}

		user, err = h.UserRepository.CreateUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: requests}
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
