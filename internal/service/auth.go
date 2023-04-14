package service

import (
	"errors"
	"maya/global"
	"maya/internal/dao"
	"maya/internal/model"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v4"
)

type authSvr struct {
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(clainsMap map[string]any) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	for k, v := range clainsMap {
		claims[k] = v
	}
	// claims["username"] = ud.Username
	// claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expire))

	t, err := token.SignedString([]byte(global.Config.Jwt.Secret))
	return t, err
}

func (a *authSvr) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := dao.Db.Where(&model.User{Email: email}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (a *authSvr) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := dao.Db.Where(&model.User{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

var (
	AuthSrv = new(authSvr)
)
