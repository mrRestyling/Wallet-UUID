package service

import "wallet/internal/models"

type Service struct {
	Storage StorageInt
}

type StorageInt interface {
	Create(wallet models.Wallet) (string, error)
	Change(wallet models.Wallet) (string, error)
	Balance(wallet models.Wallet) (string, error)
}

func New(s StorageInt) *Service {
	return &Service{
		Storage: s,
	}
}
