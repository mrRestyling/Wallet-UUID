package handlers

import (
	"log"
	"net/http"
	"strings"
	"wallet/internal/models"

	"github.com/labstack/echo"
)

const (
	// AuthorizationHeader = "Authorization"
	UserCtx = "userId"
)

// UserIdentity - middleware
func (h *Handlers) UserIdentity(c echo.Context) {
	header := c.Request().Header.Get("Authorization")

	if header == "" {
		log.Println("header is empty")
		c.JSON(http.StatusUnauthorized, "Empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		log.Println("invalid token")
		c.JSON(http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	// parse token

	userId, err := h.Serv.ParseToken(models.User{Token: headerParts[1]})

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	c.Set(UserCtx, userId)

}
