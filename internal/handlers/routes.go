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
	GenerateToken(user models.User) (string, error)
	ParseToken(user models.User) (models.User, error)
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

	// h.E.POST("/api/v1/create", h.Create)

	h.E.POST("/api/v1/wallet", h.ChangeWallet)

	h.E.GET("/api/v1/wallets/:WALLET_UUID", h.Balance)

	auth := h.E.Group("/auth")
	auth.POST("/sign-up", h.Registration)
	auth.POST("/sign-in", h.SignIn)

	api := h.E.Group("/api")
	api.POST("/v1/wallet", h.ChangeWallet)
	api.GET("/v1/wallets/:WALLET_UUID", h.Balance)

}

// TODO UserIdentity (middleware)
// Нет функции create для проверки
