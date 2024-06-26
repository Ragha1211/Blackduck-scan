package router

import (
	"order/pkg/auth"
	"order/pkg/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(orderController *handler.OrderHandler) *gin.Engine {
	router := gin.Default()

	baseRouterGroup := router.Group("/api/v1/")
	baseRouterGroup.GET("ping", handlePing)
	orderRouterGroup := baseRouterGroup.Group("/order")
	orderRouterGroup.Use(auth.ValidateJWTToken())

	controllers(orderRouterGroup, orderController)

	return router
}

func controllers(routerGroup *gin.RouterGroup, orderController *handler.OrderHandler) {
	routerGroup.GET("/", orderController.HandleGetAllOrders)
	routerGroup.GET("/:id", orderController.HandleGetOrderByID)
	routerGroup.POST("/", orderController.HandleCreateOrder)
	routerGroup.PUT("/:id", orderController.HandleUpdateOrderByID)
	routerGroup.DELETE("/:id", orderController.HandleDeleteOrderByID)
}

func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
