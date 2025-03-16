package interfaces

import (
	"context"
	"user-service/internal/models"

	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string, role string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
}

type UserService interface {
	Register(ctx context.Context, req *models.User, role string) (*models.User, error)
	Login(ctx context.Context, req models.LoginRequest, role string) (models.LoginResponse, error)
	GetProfile(ctx context.Context, username string) (models.User, error)
}

type UserAPI interface {
	RegisterUser(e echo.Context) error
	RegisterSeller(e echo.Context) error

	LoginUser(e echo.Context) error
	LoginSeller(e echo.Context) error

	GetProfile(e echo.Context) error
}
