package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handlers struct {
	E    *echo.Echo
	Serv ServInt
}

type ServInt interface {
}

func New(s ServInt) *Handlers {
	return &Handlers{
		E:    echo.New(),
		Serv: s,
	}
}

func (h *Handlers) ChangeWallet(c echo.Context) error {

	return c.JSON(http.StatusOK, "No IMP!!!") // no imp

}

func (h *Handlers) Balance(c echo.Context) error {

	return c.JSON(http.StatusOK, "No IMP!!!") // no imp

}
