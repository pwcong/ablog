package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pwcong/ablog/config"
	"github.com/pwcong/ablog/controller"
	"github.com/pwcong/ablog/errors"
)

type AuthMiddleware struct {
	Conf *config.Config
}

func (ctx *AuthMiddleware) AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Request().Cookie("TOKEN")

		var tokenString string

		if err == nil {
			tokenString = cookie.Value
		} else {
			tokenString = c.Request().Header.Get("token")
		}

		if tokenString == "" {
			return errors.NewHTTPError(http.StatusUnauthorized, controller.STATUS_ERROR, "lack of token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(ctx.Conf.Auth.Secret), nil
		})

		if err != nil {
			return errors.NewHTTPError(http.StatusUnauthorized, controller.STATUS_ERROR, "invalid token")
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			return errors.NewHTTPError(http.StatusUnauthorized, controller.STATUS_ERROR, "invalid token")
		}

		return next(c)
	}
}
