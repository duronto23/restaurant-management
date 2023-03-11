package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/odersItems", controller.GetOrderItems())
	incomingRoutes.GET("/ordersItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/ordersItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/ordersItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/ordersItems/:orderItem_id", controller.UpdateOederItem())
}
