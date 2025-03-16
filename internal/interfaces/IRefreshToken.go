package interfaces

import (
	"context"
	"user-service/helpers"
	"user-service/internal/models"

	"github.com/labstack/echo/v4"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(echo.Context) error
}
