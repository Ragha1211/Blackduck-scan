package handler

import (
	"net/http"
	"order/pkg/models/request"
	"order/pkg/models/response"
	"order/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandlerImpl(service service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}

func (handler *OrderHandler) HandleGetAllOrders(ctx *gin.Context) {

	orders := handler.orderService.FindAll()

	orderResponse := response.APIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orders,
	}
	ctx.JSON(http.StatusOK, orderResponse)
}

func (handler *OrderHandler) HandleGetOrderByID(ctx *gin.Context) {

	orderId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	order := handler.orderService.FindByID(orderId)
	orderResponse := response.APIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   order,
	}
	ctx.JSON(http.StatusOK, orderResponse)
}

func (handler *OrderHandler) HandleCreateOrder(ctx *gin.Context) {

	orderRequest := request.CreateOrderRequest{}
	err := ctx.ShouldBindJSON(&orderRequest)

	if err != nil {
		panic(err)
	}

	handler.orderService.CreateOrder(orderRequest)
	orderResponse := response.APIResponse{
		Code:   http.StatusCreated,
		Status: "Order Created",
		Data:   orderRequest,
	}
	ctx.JSON(http.StatusOK, orderResponse)
}

func (handler *OrderHandler) HandleUpdateOrderByID(ctx *gin.Context) {

	orderId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	orderRequest := request.CreateOrderRequest{}
	err = ctx.ShouldBindJSON(&orderRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	handler.orderService.UpdateOrder(orderId, orderRequest)

	orderResponse := response.APIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderRequest,
	}

	ctx.JSON(http.StatusOK, orderResponse)
}

func (handler *OrderHandler) HandleDeleteOrderByID(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	handler.orderService.DeleteOrderById(orderId)

	orderResponse := response.APIResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	ctx.JSON(http.StatusOK, orderResponse)
}
