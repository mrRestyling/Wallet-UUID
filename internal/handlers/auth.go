package handlers

import (
	"net/http"
	"wallet/internal/models"

	"github.com/labstack/echo"
)

// Registration - регистрация
func (h *Handlers) Registration(c echo.Context) error {
	const op = "internal/handlers.Registration"

	// var user models.User

	// if err := c.Bind(&user); err != nil {
	// 	log.Printf("%s: %v\n", op, err)
	// 	return c.JSON(http.StatusBadRequest, BadJSON)
	// }

	email := c.Param("email")
	password := c.Param("password")

	_, err := h.Serv.RegistrationServ(models.User{Email: email, Password: password})
	if err != nil {

	}

	return c.JSON(http.StatusOK, "!no impl")
}
