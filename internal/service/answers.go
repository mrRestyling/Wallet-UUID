package service

import "errors"

var (
	ErrWalletEmpty   = errors.New("Wallet empty")
	ErrOperationType = errors.New("Operation type empty")
	ErrAmount        = errors.New("Amount empty or negative")

	//
	IDEmpty    = "не указан ID"
	ErrIDEmpty = errors.New("id empty")

	WalletEmpty = "не указан кошелек"

	GroupEmpty = "не указана группа"

	SongNotFound = "песня не найдена"

	GoStorage = "-база данных-"
)
