package user

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	serv "github.com/hanifbg/login_register_v2/service"
	"github.com/hanifbg/login_register_v2/util/validator"
)

type service struct {
	repository   Repository
	utilPassword UtilPassword
}

func NewService(repository Repository, utilPassword UtilPassword) Service {
	return &service{
		repository,
		utilPassword,
	}
}

type CreateUserData struct {
	Name         string `validate:"required"`
	Email        string `validate:"required"`
	Phone_number string `validate:"required,number"`
	Password     string `validate:"required"`
	Address      string
}

func (s *service) CreateUser(data CreateUserData) error {
	err := validator.GetValidator().Struct(data)
	if err != nil {
		return serv.ErrInvalidData
	}

	hashedPassword, _ := s.utilPassword.EncryptPassword(data.Password)
	user := NewUser(
		data.Name,
		data.Email,
		data.Phone_number,
		string(hashedPassword),
		data.Address,
		time.Now(),
		time.Now(),
	)

	err = s.repository.CreateUser(user)
	if err != nil {
		return serv.ErrInternalServerError
	}
	return nil
}

func (s *service) LoginUser(email string, password string) (string, error) {
	userData, err := s.repository.LoginUser(email)
	if err != nil {
		return "", serv.ErrNotFound
	}

	if !s.utilPassword.ComparePassword(userData.Password, password) {
		return "", errors.New("wrong credentials")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userData.ID
	claims["role"] = "default"
	if userData.Role != 1 {
		claims["role"] = "admin"
	}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //expired token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
