package routes

import (
	controller "go-restuarant/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:order_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:order_id", controller.UpdateTable())
	
}
