// internal/service/user.go
package service

import (
	"errors"

	"github.com/CreatorQWQ/gin-admin/internal/model"
	"github.com/CreatorQWQ/gin-admin/pkg/common"
	"github.com/CreatorQWQ/gin-admin/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

var UserSvc = new(UserService)

func (s *UserService) Register(username, password, email string) error {
	var count int64
	common.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Username: username,
		Password: string(hashed),
		Email:    email,
	}

	return common.DB.Create(&user).Error
}

func (s *UserService) Login(username, password string) (string, error) {
	var user model.User
	if err := common.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
