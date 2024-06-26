package repository

import (
	"errors"
	"fmt"
	"order/pkg/models"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

// Delete implements OrderRepository.
func (o *OrderRepositoryImpl) Delete(orderId int) {
	var order models.Order
	result := o.db.Where("order_id = ?", orderId).Delete(&order)
	if result.Error != nil {
		panic(result.Error)
	}
}

// FindAll implements OrderRepository.
func (o *OrderRepositoryImpl) FindAll() (orders []models.Order) {
	result := o.db.Find(&orders)

	if result.Error != nil {
		panic(result.Error)
	}

	return

}

// FindByID implements OrderRepository.
func (o *OrderRepositoryImpl) FindByID(orderId int) (order models.Order, err error) {

	result := o.db.Find(&order, models.Order{OrderId: orderId})

	if result.Error != nil {
		err = errors.New("Order not found")
	}
	err = nil
	return
}

// Save implements OrderRepository.
func (o *OrderRepositoryImpl) Save(order models.Order) {

	result := o.db.Create(&order)
	if result.Error != nil {
		panic(result.Error)
	}
}

// Update implements OrderRepository.
func (o *OrderRepositoryImpl) Update(orderId int, order models.Order) {

	fmt.Printf("Order Id :- %v\n", orderId)

	result := o.db.Model(&models.Order{}).Where("order_id = ? ", orderId).Updates(order)
	if result.Error != nil {
		panic(result.Error)
	}

}
