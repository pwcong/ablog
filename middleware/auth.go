package middleware

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pwcong/ablog/config"
)

type AuthMiddleware struct {
	Conf *config.Config
}

func (ctx *AuthMiddleware) AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Request().Cookie("Token")

		var tokenString string
		if err == nil {
			tokenString = cookie.Value
		} else {
			tokenString = c.Request().Header.Get("Token")
		}

		if tokenString == "" {
			return errors.New("lack of token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(ctx.Conf.Auth.Secret), nil
		})

		if err != nil {
			return errors.New("invalid token")
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			return errors.New("invalid token")
		}

		return next(c)
	}
}
