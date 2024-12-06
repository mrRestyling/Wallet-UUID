package service

import (
	"errors"
	"log"
	"regexp"
	"wallet/internal/models"
)

// TODO вынести одинаковые проверки в отдельную функцию

// Create - ...
func (s *Service) Create(wallet models.Wallet) (string, error) {
	const op = "internal/service.Create"

	if wallet.WalletID == "" {
		log.Printf("%s: %v\n", op, ErrWalletEmpty)
		return "Wallet empty", ErrWalletEmpty
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$`).MatchString(wallet.WalletID) {
		log.Printf("%s: %v\n", op, ErrUUID)
		return "invalid UUID format", errors.New("invalid UUID format")
	}

	result, err := s.Storage.CreateWallet(wallet)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return result, err
	}

	return result, nil
}

// Change - ...
func (s *Service) Change(wallet models.Wallet) (string, error) {
	const op = "internal/service.Change"

	if wallet.WalletID == "" {
		log.Printf("%s: %v\n", op, ErrWalletEmpty)
		return "Wallet empty", ErrWalletEmpty
	}
	// Проверка формата UUID
	if !regexp.MustCompile(`^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$`).MatchString(wallet.WalletID) {
		log.Printf("%s: %v\n", op, ErrUUID)
		return "invalid UUID format", errors.New("invalid UUID format")
	}

	if wallet.OperationType == "" {
		log.Printf("%s: %v\n", op, ErrOperationType)

		return "Operation type empty", ErrOperationType
	}

	if wallet.OperationType != "DEPOSIT" && wallet.OperationType != "WITHDRAW" {
		log.Printf("%s: %v\n", op, ErrOperationType)
		return "Operation type empty", ErrOperationType
	}

	if wallet.Amount == 0 {
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

// Balance - ...
func (s *Service) Balance(wallet models.Wallet) (string, error) {
	const op = "internal/service.Balance"

	if wallet.WalletID == "" {
		log.Printf("%s: %v\n", op, ErrWalletEmpty)
		return "Wallet empty", errors.New("wallet ID is empty")
	}

	// Проверка формата UUID
	if !regexp.MustCompile(`^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$`).MatchString(wallet.WalletID) {
		log.Printf("%s: %v\n", op, ErrUUID)
		return "invalid UUID format", errors.New("invalid UUID format")
	}

	result, err := s.Storage.Balance(wallet)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return result, err
	}

	return result, nil

}
