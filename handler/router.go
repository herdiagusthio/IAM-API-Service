package handler

import (
	"iam-api-service/handler/user"
	"iam-api-service/middleware"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userHandler *user.Handler) {

	userV1 := e.Group("v1")
	userV1.POST("/register", userHandler.CreateUser)
	userV1.POST("/login", userHandler.LoginUser)

	cobaV1 := e.Group("v1/auth")
	cobaV1.Use(middleware.JWTMiddleware())
}
