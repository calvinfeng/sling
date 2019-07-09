package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// NewTokenAuthMiddleware returns a middleware that checks token in header.
func NewTokenAuthMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := ctx.Request().Header.Get("Token")

			if user, err := findUserByToken(db, token); err == nil {
				log.Debugf("%s from authenticated user %s", ctx.Request().Method, user.Email)
				ctx.Set("current_user", user)
				return next(ctx)
			}

			log.Errorf("%s failed with invalid token", ctx.Request().Method)
			return echo.NewHTTPError(http.StatusUnauthorized, "valid token is not presented in header")
		}
	}
}
