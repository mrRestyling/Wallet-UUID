package service

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"
	"wallet/internal/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt       = "opqwjdp234jnmj2"
	signingKey = "k34nroltnjkm2k34"
	tokenTTL   = 24 * time.Hour
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

// GenerateToken - ...
func (s *Service) GenerateToken(user models.User) (string, error) {
	const op = "internal/service.GenerateToken"

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed generate password hash")

		return "", fmt.Errorf("%s: %w", op, err)
	}
	user.PassHash = passHash

	// user.PassHash = []byte(generatePasswordHash(user.Password))

	// log.Println(user.Email, user.PassHash)

	result, err := s.Storage.GetUser(user)
	if err != nil {
		log.Println("не найден пользователь")

		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword(result.PassHash, []byte(user.Password))
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		log.Println("хэш не совпадает с присланным паролем")

		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // Токен будет не валидным через 24 часа
			IssuedAt:  time.Now().Unix(),               // Время создания токена
		},
		result.ID,
	})

	return token.SignedString([]byte(signingKey))
}

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
	// user.PassHash = []byte(generatePasswordHash(user.Password))

	log.Println(user.PassHash)

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

// ParseToken - ...
func (s *Service) ParseToken(user models.User) (models.User, error) {
	token, err := jwt.ParseWithClaims(user.Token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("метод не является HMAC")
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return models.User{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return models.User{}, errors.New("token claims are not of type")
	}

	return models.User{Claims: claims.UserID}, nil

}
