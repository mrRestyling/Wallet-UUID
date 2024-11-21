package handlers

import (
	"wallet/internal/models"

	"github.com/labstack/echo"
)

type Handlers struct {
	E    *echo.Echo
	Serv ServInt
}

type ServInt interface {
	Create(wallet models.Wallet) (string, error)
	Change(wallet models.Wallet) (string, error)
}

func New(s ServInt) *Handlers {
	return &Handlers{
		E:    echo.New(),
		Serv: s,
	}
}

func (h *Handlers) SetRoutes() {
	h.E.HideBanner = true

	h.E.POST("/api/v1/create", h.Create)

	h.E.POST("/api/v1/wallet", h.ChangeWallet)

	h.E.GET("api/v1/wallets/{WALLET_UUID}", h.Balance)

}
