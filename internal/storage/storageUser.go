package storage

import (
	"errors"
	"fmt"
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
		return Clone, ErrClone
	}

	// Создаем пользователя
	_, err = s.Db.Exec(`INSERT INTO users (email, pass_hash) VALUES ($1, $2)`, user.Email, user.PassHash)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return Internal, ErrInternal
	}

	return RegOK, nil
}

// GetUser - Получаем id, email, pass_hash для дальнейшего сравнения хеша
func (s *Storage) GetUser(user models.User) (models.User, error) {
	const op = "internal/storage.GetUser"

	var result models.User

	quary := fmt.Sprintln("SELECT id, email, pass_hash FROM users WHERE email = $1")

	smtp, err := s.Db.Prepare(quary)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return result, errors.New("Ошибка в подготовке запроса")
	}

	row := smtp.QueryRow(user.Email)

	err = row.Scan(&result.ID, &result.Email, &result.PassHash)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return result, ErrInternal
	}

	return result, err
}
