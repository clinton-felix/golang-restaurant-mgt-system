package routes

import (
	controller "github.com/clinton-felix/restaurant-mgt-project/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoice", controller.GetInvoices())
	incomingRoutes.GET("/invoice/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoice", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoice/:invoice_id", controller.UpdateInvoice())
}