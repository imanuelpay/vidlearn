package middleware

import (
	"net/http"
	"strings"
	"time"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	jwtSignedMethod = jwt.SigningMethodHS256
)

func JWTMiddleware(config *config.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Invalid token",
				})
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Invalid token",
				})
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.App.Key), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Invalid token",
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			var role float64 = claims["role"].(float64)
			if ok && token.Valid {
				if role != 1 && role != 99 {
					return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
						Message: "You dont have permission",
					})
				}
			}

			exp := claims["exp"].(float64)
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Token expired",
				})
			}

			return next(c)
		}
	}
}

func JWTMiddlewareAdmin(config *config.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Invalid token",
				})
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Invalid token",
				})
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.App.Key), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Invalid token",
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			var role float64 = claims["role"].(float64)
			if ok && token.Valid {
				if role != 99 {
					return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
						Message: "You dont have permission",
					})
				}
			}

			exp := claims["exp"].(float64)
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
					Message: "Token expired",
				})
			}

			return next(c)
		}
	}
}
