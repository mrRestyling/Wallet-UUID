package service

import "errors"

var (
	ErrWalletEmpty   = errors.New("wallet empty")
	ErrOperationType = errors.New("operation type empty")
	ErrAmount        = errors.New("amount empty or negative")
	ErrUUID          = errors.New("invalid format")

	//
	IDEmpty    = "не указан ID"
	ErrIDEmpty = errors.New("id empty")

	WalletEmpty = "не указан кошелек"

	GroupEmpty = "не указана группа"

	SongNotFound = "песня не найдена"

	GoStorage = "-база данных-"
)
