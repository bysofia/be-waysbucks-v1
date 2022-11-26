package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type ToppingRepository interface {
	FindToppings() ([]models.Topping, error)
	GetTopping(ID int) (models.Topping, error)
	AddTopping(topping models.Topping) (models.Topping, error)
	EditTopping(topping models.Topping) (models.Topping, error)
	DeleteTopping(topping models.Topping) (models.Topping, error)
}

func RepositoryTopping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindToppings() ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Preload("User").Find(&toppings).Error

	return toppings, err
}

func (r *repository) GetTopping(ID int) (models.Topping, error) {
	var topping models.Topping
	err := r.db.Preload("User").First(&topping, ID).Error

	return topping, err
}

func (r *repository) AddTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Create(&topping).Error

	return topping, err
}

func (r *repository) EditTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Save(&topping).Error

	return topping, err
}

func (r *repository) DeleteTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Delete(&topping).Error

	return topping, err

}
