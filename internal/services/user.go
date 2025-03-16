package services

import (
	"context"
	"time"
	"user-service/helpers"
	"user-service/internal/interfaces"
	"user-service/internal/models"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	UserRepo interfaces.UserRepository
}

func (r *UserServices) Register(ctx context.Context, req *models.User, role string) (*models.User, error) {
	// Generate hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)

	req.Role = role
	// Insert new user
	err = r.UserRepo.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := req
	req.Password = ""
	return resp, nil
}

func (s *UserServices) Login(ctx context.Context, req models.LoginRequest, role string) (models.LoginResponse, error) {
	var (
		response models.LoginResponse
		now      = time.Now()
	)

	// Get user by username
	userDetail, err := s.UserRepo.GetUserByUsername(ctx, req.Username, role)
	if err != nil {
		return response, errors.Wrap(err, "failed to get user by username")
	}

	// Compare password from request with password hash from database
	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return response, errors.Wrap(err, "failed to compare password")
	}

	// Generate Token
	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}

	// Generate Refresh Token
	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "refresh_token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate refresh token")
	}

	// Insert session
	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return response, errors.Wrap(err, "failed to insert new session")
	}

	response.UserID = userDetail.ID
	response.Username = userDetail.Username
	response.FullName = userDetail.FullName
	response.Email = userDetail.Email
	response.Token = token
	response.RefreshToken = refreshToken

	return response, nil
}

func (s *UserServices) GetProfile(ctx context.Context, username string) (models.User, error) {
	var (
		resp models.User 
		err error 
	)
	resp, err = s.UserRepo.GetUserByUsername(ctx, username, "")
	if err != nil {
		return resp, errors.Wrap(err, "failed to query user by username")
	}
	resp.Password = ""
	resp.Role = ""
	return resp, nil
}
