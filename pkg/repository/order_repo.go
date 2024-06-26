package repository

import "order/pkg/models"

type OrderRepository interface {
	Save(order models.Order)
	Update(orderId int, order models.Order)
	Delete(orderId int)
	FindByID(orderId int) (order models.Order, err error)
	FindAll() (orders []models.Order)
}
