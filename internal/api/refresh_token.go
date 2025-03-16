package api

import (
	"net/http"
	"user-service/constants"
	"user-service/helpers"
	"user-service/internal/interfaces"

	"github.com/labstack/echo/v4"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	refreshToken := e.Request().Header.Get("Authorization")
	token := e.Get("token")
	tokenClaim, ok := token.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse claim to claimToken")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)

	}

	resp, err := api.RefreshTokenService.RefreshToken(e.Request().Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Error("failed on RefreshToken service: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)

	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}
