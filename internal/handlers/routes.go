package handlers

func (h *Handlers) SetRoutes() {
	h.E.HideBanner = true

	h.E.POST("/api/v1/wallet", h.ChangeWallet)

	h.E.GET("api/v1/wallets/{WALLET_UUID}", h.Balance)

}
