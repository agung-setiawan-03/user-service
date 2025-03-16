package cmd

import (
	"log"
	"net/http"
	"time"
	"user-service/helpers"

	"github.com/labstack/echo/v4"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		auth := e.Request().Header.Get("authorization")
		if auth == "" {
			log.Println("Authorization empty")
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)

		}

		_, err := d.UserRepository.GetUserSessionByToken(e.Request().Context(), auth)
		if err != nil {
			log.Println("Failed to get user session on DB: ", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)

		}

		claim, err := helpers.ValidateToken(e.Request().Context(), auth)
		if err != nil {
			log.Println(err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)

		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println("jwt token is expired: ", claim.ExpiresAt)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)

		}

		e.Set("token", claim)

		return next(e)
	}

}
