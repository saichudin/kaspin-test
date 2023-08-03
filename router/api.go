package router

import (
	paymentHandler "saichudin/kaspin/payment"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	payment := e.Group("/nicepay")
	payment.POST("/registration", paymentHandler.Register)
	payment.POST("/inquiry", paymentHandler.Inquiry)
	payment.POST("/payment", paymentHandler.Payment)
}
