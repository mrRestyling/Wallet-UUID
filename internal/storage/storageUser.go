package storage

import (
	"log"
	"wallet/internal/models"
)

// CreateUser - ...
func (s *Storage) CreateUser(user models.User) (string, error) {
	const op = "internal/storage.CreateUser"

	// Создаем пользователя
	_, err := s.Db.Exec(`INSERT INTO users (wallet_id, email, pass_hash) VALUES ($1, $2, $3)`, user.WalletID, user.Email, user.PassHash)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return Internal, ErrInternal
	}

	return AddOK, nil

}
