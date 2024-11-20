package main

import (
	"wallet/internal/handlers"
)

func main() {

	h := handlers.New()

	h.E.POST("/api/v1/wallet", h.ChangeWallet)

	h.E.Start(":8080")

}
