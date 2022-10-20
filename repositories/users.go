package repositories

import (
	"counting_discount/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	FindProductById(ProductID []int) ([]models.Product, error)
	GetUser(ID int) (models.User, error)
	CreateUser(User models.User) (models.User, error)
	UpdateUser(User models.User) (models.User, error)
	DeleteUser(User models.User) (models.User, error)
}

func RepositoryUsers(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Product").Find(&users).Error

	return users, err
}

func (r *repository) FindProductById(ProductID []int) ([]models.Product, error) {
	var product []models.Product
	err := r.db.Find(&product, ProductID).Error

	return product, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Product").First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	r.db.Model(&user).Association("Product").Replace(user.Product)

	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
