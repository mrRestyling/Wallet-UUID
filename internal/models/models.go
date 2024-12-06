package models

import "time"

type Wallet struct {
	ID uint `json:"id"`

	WalletID      string `json:"valletId"`
	OperationType string `json:"operationType"`
	Amount        uint   `json:"amount"`

	Balance     uint      `json:"balance"`
	CreatedTime time.Time `json:"create"`
	UpdatedTime time.Time `json:"update"`
}

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	PassHash []byte `json:"passHash"`
	WalletID string `json:"walletId"`
}
