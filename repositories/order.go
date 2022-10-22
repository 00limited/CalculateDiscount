package repositories

import (
	"counting_discount/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	FindUserById(UserID []int) ([]models.User, error)
	CreateOrder(Order models.Order) (models.Order, error)
	DeleteOrder(Order models.Order) (models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.Order, error) {
	var orders []models.Order

	err := r.db.Preload("User").Find(&orders).Error

	return orders, err
}
func (r *repository) FindUserById(UserID []int) ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user, UserID).Error

	return user, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("User").First(&order, ID).Error

	return order, err
}

func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Preload("User").Create(&order).Error

	return order, err
}

func (r *repository) DeleteOrder(order models.Order) (models.Order, error) {
	err := r.db.Delete(&order).Error

	return order, err
}
