package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handlers struct {
	E *echo.Echo
}

func New() *Handlers {
	return &Handlers{
		E: echo.New(),
	}
}

func (h *Handlers) ChangeWallet(c echo.Context) error {

	return c.JSON(http.StatusOK, "No IMP!!!")

}

func (h *Handlers) CheckWallet() {

}
