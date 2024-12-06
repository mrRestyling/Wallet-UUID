package service

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"wallet/internal/models"

	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// RegistrationServ - ...
func (s *Service) RegistrationServ(user models.User) (string, error) {
	const op = "internal/service.RegistrationServ"

	if user.Email == "" || !isValidEmail(user.Email) {
		log.Println(user.Email)
		log.Println("err Email")
		return "", errors.New("invalid format")
	}

	if user.Password == "" {
		log.Println("err Password")
		return "", errors.New("invalid format")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed generate password hash")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	user.PassHash = passHash

	// user.WalletID = uuid.New().String()

	id, err := s.Storage.CreateUser(user)
	if err != nil {
		log.Println("failed create user")

		return id, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
