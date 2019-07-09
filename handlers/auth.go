package handlers

import (
	"net/http"

	"github.com/calvinfeng/go-academy/userauth/model"
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
		c := &Credential{}
		if err := ctx.Bind(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByCredentials(db, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}

		return ctx.JSON(http.StatusOK, TokenResponse{
			Name:     user.Name,
			Email:    user.Email,
			JWTToken: user.JWTToken,
		})
	}

}

// GetCurrentUserHandler returns a handler that gets the current user.
func GetCurrentUserHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, ctx.Get("current_user"))
	}
}

// GetUsersHandler returns a handler that gets all current users.
func GetUsersHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var users []*model.User

		if err := db.Find(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		// Clear token for security
		for _, u := range users {
			u.JWTToken = ""
		}

		return ctx.JSON(http.StatusOK, users)
	}
}
