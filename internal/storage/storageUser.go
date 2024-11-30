package storage

import (
	"log"
	"wallet/internal/models"
)

// CreateUser - ...
func (s *Storage) CreateUser(user models.User) (string, error) {
	const op = "internal/storage.CreateUser"

	// Проверяем существование пользователя
	var uID int64
	err := s.Db.Get(&uID, "SELECT id FROM users WHERE email = $1", user.Email)
	if err == nil {
		log.Println("Пользователь уже существует")
	}

	// Создаем пользователя
	_, err = s.Db.Exec(`INSERT INTO users (email, pass_hash) VALUES ($1, $2)`, user.Email, user.PassHash)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return Internal, ErrInternal
	}

	return RegOK, nil
}
