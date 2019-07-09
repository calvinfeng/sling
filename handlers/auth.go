package handlers

import (
	"net/http"

	"github.com/jchou8/sling/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type (
	// Credential is a payload that captures user submitted credentials.
	Credential struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// TokenResponse is a payload that returns JWT token back to client.
	TokenResponse struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		JWTToken string `json:"jwt_token"`
	}
)

// NewUserHandler returns a handler that creates a new user.
func NewUserHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := &models.User{}
		if err := ctx.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if err := user.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
		}

		hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user.Password = ""
		user.PasswordDigest = hashBytes

		if err := db.Create(user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, user)
	}
}

// LoginHandler returns a handler that logs a user in.
func LoginHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//TODO: Implement
		return nil
	}

}

// GetCurrentUserHandler returns a handler that gets the current user.
func GetCurrentUserHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//TODO: Implement
		return nil
	}
}

// GetUsersHandler returns a handler that gets all current users.
func GetUsersHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//TODO: Implement
		return nil
	}
}
