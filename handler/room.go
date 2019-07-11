package handler

import (
	"net/http"

	"github.com/calvinfeng/sling/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// GetRoomsHandler returns a handler that gets all current rooms.
func GetRoomsHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Grab current user
		token := ctx.Request().Header.Get("Token")
		user, err := findUserByToken(db, token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		rooms, dbErr := model.GetRooms(db, user.ID)
		if dbErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, rooms)
	}
}
