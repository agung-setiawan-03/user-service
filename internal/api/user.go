package api

import (
	"net/http"
	"user-service/constants"
	"user-service/helpers"
	"user-service/internal/interfaces"
	"user-service/internal/models"

	"github.com/labstack/echo/v4"
)

type UserAPI struct {
	UserService interfaces.UserService
}

func (api *UserAPI) RegisterUser(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	// Parse first request from client
	if err := e.Bind(&req); err != nil {
		log.Error("Failed to parse request", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}

	// Validate request 
	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}


	// Register new user
	resp, err := api.UserService.RegisterUser(e.Request().Context(), &req)
	if err != nil {
		log.Error("Failed to register new user", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}


func (api *UserAPI) RegisterSeller(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	// Parse first request from client
	if err := e.Bind(&req); err != nil {
		log.Error("Failed to parse request", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}

	// Validate request 
	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}


	// Register new user
	resp, err := api.UserService.RegisterSeller(e.Request().Context(), &req)
	if err != nil {
		log.Error("Failed to register new user", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}
