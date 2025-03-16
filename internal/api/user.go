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
		log.Error("Failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}

	// Validate request 
	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}


	// Register new user
	resp, err := api.UserService.Register(e.Request().Context(), &req, "user")
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
		log.Error("Failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}

	// Validate request 
	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}


	// Register new seller
	resp, err := api.UserService.Register(e.Request().Context(), &req, "seller")
	if err != nil {
		log.Error("Failed to register new seller: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}


func (api *UserAPI) LoginUser(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.LoginRequest{}

	// Parse first request from client
	if err := e.Bind(&req); err != nil {
		log.Error("Failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}

	// Validate request 
	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}


	// Login user
	resp, err := api.UserService.Login(e.Request().Context(), req, "user")
	if err != nil {
		log.Error("Failed to login user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}



func (api *UserAPI) LoginSeller(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.LoginRequest{}

	// Parse first request from client
	if err := e.Bind(&req); err != nil {
		log.Error("Failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}

	// Validate request 
	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrBadrequest, nil)
	}


	// Register seller
	resp, err := api.UserService.Login(e.Request().Context(), req, "seller")
	if err != nil {
		log.Error("Failed to login user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}

func (api *UserAPI) GetProfile(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	token := e.Get("token")
	tokenClaim, ok := token.(*helpers.ClaimToken)
	if ! ok {
		log.Error("Failed to fetch token")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)
	}
	
	// GetProfile User 
	resp, err := api.UserService.GetProfile(e.Request().Context(), tokenClaim.Username)
	if err != nil {
		log.Error("Failed to get profile: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrInternalServer, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}


