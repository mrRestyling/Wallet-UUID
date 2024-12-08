package handlers

import (
	"log"
	"net/http"
	"wallet/internal/models"

	"github.com/labstack/echo"
)

// Registration - регистрация
func (h *Handlers) Registration(c echo.Context) error {
	const op = "internal/handlers.Registration"

	var user models.User

	if err := c.Bind(&user); err != nil {
		log.Printf("%s: %v\n", op, err)
		return c.JSON(http.StatusBadRequest, BadJSON)
	}

	log.Println(user)

	result, err := h.Serv.RegistrationServ(models.User{
		Email:    user.Email,
		Password: user.Password})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

// SignIn - авторизация
func (h *Handlers) SignIn(c echo.Context) error {
	const op = "internal/handlers.Login"

	var user models.User

	if err := c.Bind(&user); err != nil {
		log.Printf("%s: %v\n", op, err)
		return c.JSON(http.StatusBadRequest, BadJSON)
	}

	log.Println(user)

	token, err := h.Serv.GenerateToken(models.User{
		Email:    user.Email,
		Password: user.Password})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
