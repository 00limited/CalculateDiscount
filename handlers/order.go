package handlers

import (
	orderdto "counting_discount/dto/Order"
	dto "counting_discount/dto/result"
	"counting_discount/models"
	"counting_discount/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
	UserRepository  repositories.UserRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository, UserRepository repositories.UserRepository) *handlerOrder {
	return &handlerOrder{OrderRepository, UserRepository}
}

func (h *handlerOrder) FindOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.OrderRepository.FindOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: orders}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerOrder) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var order models.Order
	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseOrder(order)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := orderdto.OrderRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// user, _ := h.OrderRepository.FindUserById(request.UserId)

	//counting all price for all buyer
	var totalPrice int
	for i := 1; i <= len(request.UserId); i++ {
		pddetail, _ := h.UserRepository.GetUser(i)
		totalPrice += pddetail.Total
	}

	var discountUser int
	for i := 1; i <= len(request.UserId); i++ {
		pddetail, _ := h.UserRepository.GetUser(i)
		discountUser = (pddetail.Total / totalPrice) * 100
	}

	//counting price after
	var countDis int
	if (totalPrice*request.Discount)/100 > request.Maxdiscount {
		countDis += request.Maxdiscount
	} else {
		countDis = (totalPrice * request.Discount) / 100
	}

	//counting final price after discount
	finalPrice := totalPrice - countDis

	//counting final price for each user

	for i := 1; i <= len(request.UserId); i++ {
		pddetail, _ := h.UserRepository.GetUser(i)
		pddetail.Total = (finalPrice * discountUser) / 100
		_, _ = h.UserRepository.UpdateTotal(pddetail, request.UserId[i])
		// user = append(user, newUser)
	}

	// for i

	// order := models.Order{
	// 	Discount: request.Discount,
	// 	Total:    finalPrice,
	// 	MaxDiscount: request.Maxdiscount,
	// 	Users: 	user,
	// }

	// order, err = h.OrderRepository.CreateOrder(order)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// order, _ = h.OrderRepository.GetOrder(order.ID)

	w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Code: http.StatusOK, Data: order}
	// json.NewEncoder(w).Encode(response)
}

func convertResponseOrder(p models.Order) models.Order {
	return models.Order{
		ID:       p.ID,
		Discount: p.Discount,
		Total:    p.Total,
	}
}
