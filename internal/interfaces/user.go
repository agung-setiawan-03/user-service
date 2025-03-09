package interfaces

import (
	"context"
	"user-service/internal/models"

	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}

type UserService interface {
	RegisterUser(ctx context.Context, req *models.User) (*models.User, error)
	RegisterSeller(ctx context.Context, req *models.User) (*models.User, error)
}

type UserAPI interface {
	RegisterUser(e echo.Context) error 
	RegisterSeller(e echo.Context) error
}
