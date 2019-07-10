package handler

import (
	"github.com/calvinfeng/sling/model"
	"github.com/calvinfeng/sling/util"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func findUserByToken(db *gorm.DB, token string) (*model.User, error) {
	var user model.User
	if err := db.Where("jwt_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func findUserByCredentials(db *gorm.DB, c *Credential) (*model.User, error) {
	var user model.User
	if err := db.Where("name = ?", c.Username).First(&user).Error; err != nil {
		util.LogErr("fail on fetch", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(c.Password)); err != nil {
		return nil, err
	}

	util.LogInfo("success on password")

	return &user, nil
}
