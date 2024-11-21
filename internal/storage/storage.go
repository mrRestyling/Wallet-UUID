package storage

import (
	"fmt"
	"log"
	"strconv"
	"wallet/internal/models"
)

func (s *Storage) Create(wallet models.Wallet) (string, error) {
	const op = "internal/storage.Create"

	// Проверяем существование кошелека
	var wID int64
	err := s.Db.Get(&wID, "SELECT id FROM wallets WHERE wallet_id = $1", wallet.WalletID)
	if err == nil {
		log.Println("Кошелек уже существует")

		return strconv.Itoa(int(wID)), ErrClone
	}

	// Создаем кошелек
	_, err = s.Db.Exec(`INSERT INTO wallets (wallet_id, balance) VALUES ($1, $2)`, wallet.WalletID, wallet.Balance)
	if err != nil {
		fmt.Printf("%s: %v\n", op, err)
		return Internal, ErrInternal
	}

	return AddOK, nil
}

func (s *Storage) Change(wallet models.Wallet) (string, error) {
	const op = "internal/storage.Change"

	// Проверяем существование кошелека
	var wID int64
	err := s.Db.Get(&wID, "SELECT id FROM wallets WHERE wallet_id = $1", wallet.WalletID)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		log.Println("Кошелек не найден")
		return "Кошелек не найден", ErrNotFound
	}

	var query string

	if wallet.OperationType == "DEPOSIT" {
		query = `UPDATE wallets SET balance = balance + $1, updated_at = NOW() WHERE id = $2`
	} else {
		query = `UPDATE wallets SET balance = balance - $1, updated_at = NOW() WHERE id = $2`
	}

	// Обновляем баланс кошелька
	_, err = s.Db.Exec(query, wallet.Amount, wID)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return "Ошибка обновления баланса", ErrInternal
	}

	return AddOK, nil

}

func (s *Storage) Balance(wallet models.Wallet) (string, error) {
	const op = "internal/storage.Balance"

	var balance int64

	err := s.Db.Get(&balance, "SELECT balance FROM wallets WHERE wallet_id = $1", wallet.WalletID)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		log.Println("Кошелек не найден")
		return "Кошелек не найден", ErrNotFound
	}

	return strconv.Itoa(int(balance)), nil

}
