package service

import (
	"log"
	"wallet/internal/models"
)

func (s *Service) Create(wallet models.Wallet) (string, error) {
	const op = "internal/service.Create"

	if wallet.WalletID == "" {
		log.Printf("%s: %v\n", op, ErrWalletEmpty)
		return "Wallet empty", ErrWalletEmpty
	}

	result, err := s.Storage.Create(wallet)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		log.Println("TUTA", err)
		return result, err
	}

	return "", nil
}

func (s *Service) Change(wallet models.Wallet) (string, error) {
	const op = "internal/service.Change"

	if wallet.WalletID == "" {
		log.Printf("%s: %v\n", op, ErrWalletEmpty)
		return "Wallet empty", ErrWalletEmpty
	}

	if wallet.OperationType == "" || wallet.OperationType == "DEPOSIT" || wallet.OperationType == "WITHDRAW" {
		log.Printf("%s: %v\n", op, ErrOperationType)

		return "Operation type empty", ErrOperationType
	}

	if wallet.Amount == 0 || wallet.Amount < 0 {
		log.Printf("%s: %v\n", op, ErrAmount)

		return "Amount empty or negative", ErrAmount
	}

	result, err := s.Storage.Change(wallet)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return result, err
	}

	return result, nil
}
