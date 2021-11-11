package user

import (
	"iam-api-service/common"
	"iam-api-service/handler/user/request"
	"iam-api-service/service/user"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) CreateUser(c echo.Context) error {
	createUserReq := new(request.CreateUserRequest)

	if err := c.Bind(createUserReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.CreateUser(*createUserReq.ConvertToUserData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (handler *Handler) LoginUser(c echo.Context) error {
	createLoginReq := new(request.LoginUserRequest)

	if err := c.Bind(createLoginReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	token, err := handler.service.LoginUser(createLoginReq.Email, createLoginReq.Password)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(token))
}

func (handler *Handler) CreateAdmin(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		return c.JSON(common.NewForbiddenResponse())
	}

	createUserReq := new(request.CreateUserRequest)

	if err := c.Bind(createUserReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.CreateAdmin(*createUserReq.ConvertToUserData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
