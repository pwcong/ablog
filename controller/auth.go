package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type AuthForm struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type AuthController struct {
	Base *BaseController
}

func (ctx *AuthController) Login(c echo.Context) error {

	ctx.Base.Service.Info(c.RealIP(), "login")

	uf := new(AuthForm)
	if err := c.Bind(uf); err != nil {
		return err
	}

	if uf.Username == "" || uf.Password == "" {
		return errors.New("username and password can not be empty")
	}

	if uf.Username != ctx.Base.Conf.Auth.Username || uf.Password != ctx.Base.Conf.Auth.Password {
		return errors.New("invalid username or password")
	}

	now := time.Now()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expiredTime := ctx.Base.Conf.Auth.ExpiredTime
	claims["exp"] = now.Add(time.Second * time.Duration(expiredTime)).Unix()

	t, err := token.SignedString([]byte(ctx.Base.Conf.Auth.Secret))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.Name = "TOKEN"
	cookie.Value = t
	cookie.Expires = now.Add(time.Duration(expiredTime/3600) * time.Hour)
	c.SetCookie(cookie)

	return BaseResponse(c, true, STATUS_OK, "login successfully", struct {
		Token string `json:"token"`
	}{
		Token: t,
	})

}
