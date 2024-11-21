package handlers

import (
	"log"
	"net/http"
	"wallet/internal/models"

	"github.com/labstack/echo"
)

func (h *Handlers) Create(c echo.Context) error {
	const op = "internal/handlers.Create"

	var wallet models.Wallet

	if err := c.Bind(&wallet); err != nil {
		log.Println("BadJSON", err)
		log.Printf("%s: %v\n", op, err)
		return c.JSON(http.StatusBadRequest, BadJSON)
	}

	// Service
	result, err := h.Serv.Create(wallet)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return h.ModelError(c, err, result)
	}

	log.Printf("%s: %s\n", op, "Success")
	return c.JSON(http.StatusOK, result)
}

func (h *Handlers) ChangeWallet(c echo.Context) error {
	const op = "internal/handlers.ChangeWallet"

	var wallet models.Wallet

	if err := c.Bind(&wallet); err != nil {
		log.Println("BadJSON", err)
		log.Printf("%s: %v\n", op, err)
		return c.JSON(http.StatusBadRequest, BadJSON)
	}

	// Service
	result, err := h.Serv.Change(wallet)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return h.ModelError(c, err, result)
	}

	log.Printf("%s: %s\n", op, "Success")
	return c.JSON(http.StatusOK, result)
}

func (h *Handlers) Balance(c echo.Context) error {
	const op = "internal/handlers.Balance"

	walletID := c.Param("WALLET_UUID")

	result, err := h.Serv.Balance(models.Wallet{WalletID: walletID})
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return h.ModelError(c, err, result)
	}

	log.Printf("%s: %s\n", op, "Success")
	return c.JSON(http.StatusOK, result)
}
