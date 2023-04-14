package service

import (
	"errors"
	"maya/internal/dao"
	"maya/internal/model"
	"strconv"

	// jwtv4 "github.com/gofiber/jwt/v3"
	jwtv4 "github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidToken(t *jwtv4.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwtv4.MapClaims)
	uid := int(claims["userid"].(float64))

	return uid == n
}

func ValidUser(id string, p string) bool {
	var user model.User

	dao.Db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

type userSrv struct {
}

func (usrv *userSrv) GetUser(id int) (*model.User, error) {
	var user model.User
	dao.Db.Find(&user, id)
	if user.Username == "" {
		return nil, errors.New("没找到指定id的用户信息")
	}
	return &user, nil
}

func (usrv *userSrv) CreateUser(u *model.User) error {
	dao.Db.Create(u)
	if err := dao.Db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (usrv *userSrv) UpdateUser(u *model.User) error {
	var user model.User

	dao.Db.First(&user, u.ID)
	user.Address = u.Address

	if err := dao.Db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (usrv *userSrv) DeleteUser(u *model.User) error {
	var user model.User

	dao.Db.First(&user, u.ID)

	if err := dao.Db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// 为了将服务细化到具体类里，用公共变量导出User的方法，让调用者方便找到对应的方法
var (
	UserSrv = new(userSrv)
)
