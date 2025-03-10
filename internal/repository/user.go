package repository

import (
	"context"
	"errors"
	"user-service/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string, role string) (models.User, error) {
	var (
		user models.User
		err  error
	)

	sql := r.DB.Where("username = ?", username) // SELECT * FROM users WHERE username = ?

	if role != "" {
		sql = sql.Where("role = ?", role) // SELECT * FROM users WHERE username = ? AND role = ?
	}

	err = sql.First(&user).Error // SELECT * FROM users WHERE username = ? AND role = ? LIMIT 1
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(session).Error
}