package service

import (
	"order/pkg/models/request"
	"order/pkg/models/response"
)

type OrderService interface {
	FindByID(orderId int) response.OrderResponse
	FindAll() []response.OrderResponse
	CreateOrder(req request.CreateOrderRequest)
	UpdateOrder(orderId int, req request.CreateOrderRequest)
	DeleteOrderById(orderId int)
}
