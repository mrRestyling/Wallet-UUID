package service

import "wallet/internal/models"

type Service struct {
	Storage StorageItn
}

type StorageItn interface {
	Create(wallet models.Wallet) (string, error)
	Change(wallet models.Wallet) (string, error)
}

func New(s StorageItn) *Service {
	return &Service{
		Storage: s,
	}
}
