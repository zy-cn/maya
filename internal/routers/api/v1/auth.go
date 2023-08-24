package v1

import (
	"maya/internal/model"
	"maya/internal/routers"
	srv "maya/internal/service"
	"maya/pkg/errcode"
	"net/mail"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginBody struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	loginBody := new(LoginBody)
	r := routers.NewResResult[any]()

	if err := c.BodyParser(&loginBody); err != nil {
		// r.Code = errcode.InvalidParam
		// r.Message = errcode.GetDesc(r.Code)
		r.SetErrorCode(errcode.InvalidParam)
		r.Data = nil
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	identity := loginBody.Identity
	pass := loginBody.Password
	var user *model.User
	var err error

	_, err = mail.ParseAddress(identity)
	if err == nil { //如果是email,按email查询
		user, err = srv.AuthSrv.GetUserByEmail(identity)
		if err != nil {
			// r.Code = errcode.DbQueryFailed
			// r.Message = errcode.GetDesc(r.Code)
			r.SetErrorCode(errcode.DbQueryFailed)
			r.Data = err
			return c.Status(fiber.StatusUnauthorized).JSON(r)
		}
	} else {
		user, err = srv.AuthSrv.GetUserByUsername(identity)
		if err != nil {
			// r.Code = errcode.DbQueryFailed
			// r.Message = errcode.GetDesc(r.Code)
			r.SetErrorCode(errcode.DbQueryFailed)
			r.Data = err
			return c.Status(fiber.StatusUnauthorized).JSON(r)
		}
	}

	if user == nil {
		// r.Code = errcode.DbNotExists
		// r.Message = errcode.GetDesc(r.Code)
		r.SetErrorCode(errcode.DbNotExists)
		r.Data = nil
		return c.Status(fiber.StatusUnauthorized).JSON(r)
	}

	if !srv.CheckPasswordHash(pass, user.Password) {
		// r.Code = errcode.InvalidPassword
		// r.Message = errcode.GetDesc(r.Code)
		r.SetErrorCode(errcode.InvalidPassword)
		r.Data = nil
		return c.Status(fiber.StatusUnauthorized).JSON(r)
	}

	claims := make(map[string]any)
	claims["username"] = user.Username
	claims["userid"] = user.ID

	token, err := srv.GenerateToken(claims)
	if err != nil {
		// r.Code = -1
		// r.Message = "生成token失败"
		r.SetErrorCode(errcode.GenerateTokenFail)
		r.Data = nil
		return c.Status(fiber.StatusUnauthorized).JSON(r)
	}

	r.Code = errcode.Success
	r.Data = token

	return c.JSON(r)
}
