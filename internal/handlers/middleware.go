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
func (h *Handlers) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")

		if header == "" {
			log.Println("header is empty")
			c.JSON(http.StatusUnauthorized, "Empty authorization header")
			return echo.NewHTTPError(http.StatusUnauthorized, "Empty authorization header")
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 {
			log.Println("invalid token")
			c.JSON(http.StatusUnauthorized, "Invalid authorization header")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header")
		}

		// parse token

		userID, err := h.Serv.ParseToken(models.User{Token: headerParts[1]})

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, "Invalid authorization header")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header")
		}

		c.Set(UserCtx, userID)
		log.Println(userID)

		return next(c)
	}

}
