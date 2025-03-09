package services

import (
	"context"
	"user-service/internal/interfaces"
	"user-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	UserRepo interfaces.UserRepository
}

func (r *UserServices) RegisterUser(ctx context.Context, req *models.User) (*models.User, error) {
	// Generate hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)

	req.Role = "user"
	// Insert new user 
	err = r.UserRepo.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}
	
	resp := req 
	req.Password = ""
	return resp, nil
}

func (r *UserServices) RegisterSeller(ctx context.Context, req *models.User) (*models.User, error) {
	// Generate hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)

	req.Role = "seller"
	// Insert new user 
	err = r.UserRepo.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}
	
	resp := req 
	req.Password = ""
	return resp, nil
}
