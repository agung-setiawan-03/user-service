package cmd

import (
	"user-service/helpers"
	"user-service/internal/api"
	"user-service/internal/interfaces"
	"user-service/internal/repository"
	"user-service/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInjection()
	e := echo.New()
	e.GET("/healthcheck", d.Healthcheck.HealthCheck)
	userV1 := e.Group("user/v1")
	userV1.POST("/register", d.UserAPI.RegisterUser)
	userV1.POST("/register/seller", d.UserAPI.RegisterSeller)
	userV1.POST("/login", d.UserAPI.LoginUser)
	userV1.POST("/login/seller", d.UserAPI.LoginSeller)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	Healthcheck *api.HealthCheckAPI
	UserAPI     interfaces.UserAPI
}

func dependencyInjection() Dependency {
	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}
	userSvc := &services.UserServices{
		UserRepo: userRepo,
	}
	userAPI := &api.UserAPI{
		UserService: userSvc,
	}

	return Dependency{
		Healthcheck: &api.HealthCheckAPI{},
		UserAPI:     userAPI,
	}
}
