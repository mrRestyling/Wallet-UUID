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

	return AddGroupOK, nil
}

func (s *Storage) Change(wallet models.Wallet) (string, error) {
	const op = "internal/storage.Change"

	// Проверяем существование кошелека
	var wID int64
	err := s.Db.Get(&wID, "SELECT id FROM wallets WHERE wallet_id = $1", wallet.WalletID)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		log.Println("Кошелек не найден")
		return "Кошелек не существует", ErrNotFound
	}

	return "", nil // no imp

}
