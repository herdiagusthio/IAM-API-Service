package user

import (
	"fmt"
	"time"

	"iam-api-service/service/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"name; type:varchar(50); not null"`
	Email       string `gorm:"email; type:varchar(50); unique; not null"`
	PhoneNumber string `gorm:"phone_number; unique; not null"`
	Password    string `gorm:"password; not null"`
	RoleID      uint   `gorm:"role_id; not null"`
}

type Role struct {
	gorm.Model
	RoleName string `gorm:"name; not null"`
}

type GormRepository struct {
	DB *gorm.DB
}

func newUserTable(user user.User) *User {

	return &User{
		gorm.Model{user.ID, user.CreatedAt, user.UpdatedAt, gorm.DeletedAt{time.Time{}, false}},
		user.Name,
		user.Email,
		user.Phone_number,
		user.Password,
		user.RoleID,
	}
}
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (col *User) ToUserService() user.User {
	var user user.User

	user.ID = col.ID
	user.CreatedAt = col.CreatedAt
	user.UpdatedAt = col.UpdatedAt
	user.DeletedAt = &col.DeletedAt.Time
	user.Name = col.Name
	user.Email = col.Email
	user.Phone_number = col.PhoneNumber
	user.Password = col.Password
	user.RoleID = col.RoleID

	return user
}

func (repo *GormRepository) CreateUser(user user.User) error {
	userData := newUserTable(user)

	err := repo.DB.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormRepository) LoginUser(email string) (*user.User, error) {
	var userData User

	err := repo.DB.Where("email = ?", email).First(&userData).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	user := userData.ToUserService()

	return &user, nil
}
