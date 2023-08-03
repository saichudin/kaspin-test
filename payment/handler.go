package payment

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	form := new(RegistrationForm)
	if err := c.Bind(form); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	data, msg := RegistrationService(form)

	resp := DefaultResponse{
		Message: msg,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func Inquiry(c echo.Context) error {
	form := new(InquiryForm)
	if err := c.Bind(form); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	data, msg := InquiryService(form)

	resp := DefaultResponse{
		Message: msg,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func Payment(c echo.Context) error {
	resp := DefaultResponse{
		Message: "Payment Success",
		Data:    "",
	}
	return c.JSON(http.StatusOK, resp)
}
