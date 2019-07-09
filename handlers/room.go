package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// GetRoomsHandler returns a handler that gets all current rooms.
func GetRoomsHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//TODO: Determine room model
		/*var rooms []*models.Room

		if err := db.Find(&rooms).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		return ctx.JSON(http.StatusOK, rooms)*/
	}
}
