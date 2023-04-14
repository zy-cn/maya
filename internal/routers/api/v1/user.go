package v1

import (
	"maya/internal/model"
	"maya/internal/routers"
	srv "maya/internal/service"
	"maya/pkg/errcode"
	"strconv"

	jwtv4 "github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	r := routers.NewResResult[any]()

	var userInput NewUser
	if err := c.BodyParser(&userInput); err != nil {
		r.Code = errcode.InvalidParam
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)

	}

	hash, err := srv.HashPassword(userInput.Password)
	if err != nil {
		r.Code = errcode.HashError
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}

	user := new(model.User)
	user.Password = hash
	if err := srv.UserSrv.CreateUser(user); err != nil {
		r.Code = errcode.DbCreatFailed
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}

	r.Data = user

	return c.JSON(r)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	r := routers.NewResResult[any]()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		r.Code = errcode.InvalidParam
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}
	user, err := srv.UserSrv.GetUser(idInt)
	if err != nil {
		r.Code = errcode.DbNotExists
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}

	r.Code = errcode.Success
	r.Data = *user
	return c.JSON(r)
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	r := routers.NewResResult[any]()

	user := new(model.User)

	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		r.Code = errcode.InvalidParam
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)

	}

	id := c.Params("id")
	token := c.Locals("user").(*jwtv4.Token)
	if !srv.ValidToken(token, id) {
		r.Code = errcode.InvalidToken
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}

	if err := srv.UserSrv.UpdateUser(user); err != nil {
		r.Code = errcode.DbUpdateFailed
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}

	r.Data = user

	return c.JSON(r)
}

func DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	r := routers.NewResResult[any]()
	if err := c.BodyParser(&pi); err != nil {
		r.Code = errcode.InvalidParam
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwtv4.Token)

	if !srv.ValidToken(token, id) {
		r.Code = errcode.InvalidToken
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)
	}

	if !srv.ValidUser(id, pi.Password) {
		r.Code = errcode.InvalidPassword
		r.Message = errcode.GetDesc(r.Code)
		return c.Status(500).JSON(r)

	}

	idInt, _ := strconv.Atoi(id)
	user := new(model.User)
	user.ID = uint32(idInt)
	srv.UserSrv.DeleteUser(user)

	return c.JSON(r)
}
