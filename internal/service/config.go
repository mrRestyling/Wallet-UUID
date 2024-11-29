package service

import "wallet/internal/models"

// Service - ...
type Service struct {
	Storage StorageInt
}

// StorageInt - ...
type StorageInt interface {
	CreateWallet(wallet models.Wallet) (string, error)
	Change(wallet models.Wallet) (string, error)
	Balance(wallet models.Wallet) (string, error)

	CreateUser(user models.User) (string, error)
}

// New - ...
func New(s StorageInt) *Service {
	return &Service{
		Storage: s,
	}
}
