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
		var rooms []*model.Room
		// need to query all rooms, mark registered, joined and has notification
		if rooms, err := db.Find(&rooms).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		return ctx.JSON(http.StatusOK, rooms)
	}
}
