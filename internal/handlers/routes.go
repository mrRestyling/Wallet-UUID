package handlers

import (
	"wallet/internal/models"

	"github.com/labstack/echo"
)

// Handlers - ...
type Handlers struct {
	E    *echo.Echo
	Serv ServInt
}

// ServInt - ...
type ServInt interface {
	Create(wallet models.Wallet) (string, error)
	Change(wallet models.Wallet) (string, error)
	Balance(wallet models.Wallet) (string, error)
	RegistrationServ(user models.User) (string, error)
}

// New - ...
func New(s ServInt) *Handlers {
	return &Handlers{
		E:    echo.New(),
		Serv: s,
	}
}

// SetRoutes - ...
func (h *Handlers) SetRoutes() {
	h.E.HideBanner = true

	// h.E.POST("/login", h.Login)

	h.E.POST("/registration", h.Registration)

	// h.E.POST("/api/v1/create", h.Create)

	h.E.POST("/api/v1/wallet", h.ChangeWallet)

	h.E.GET("/api/v1/wallets/:WALLET_UUID", h.Balance)

}
