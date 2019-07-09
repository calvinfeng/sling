package handlers

import (
	"github.com/jchou8/sling/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func findUserByToken(db *gorm.DB, token string) (*models.User, error) {
	var user models.User
	if err := db.Where("jwt_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func findUserByCredentials(db *gorm.DB, c *Credential) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", c.Username).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(c.Password)); err != nil {
		return nil, err
	}

	return &user, nil
}
