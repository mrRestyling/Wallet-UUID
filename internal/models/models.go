package models

import "time"

type Wallet struct {
	ID uint `json:"id"`

	WalletID      string `json:"valletId"`
	OperationType string `json:"operationType"`
	Amount        uint    `json:"amount"`

	Balance     uint      `json:"balance"`
	CreatedTime time.Time `json:"create"`
	UpdatedTime time.Time `json:"update"`
}
