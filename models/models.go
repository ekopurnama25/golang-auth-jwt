package models

import (
	"golang.org/x/crypto/bcrypt"
)
type Users struct {
	UserId 			uint  `gorm:"primaryKey" json:"user_id"`
	Username       string `gorm:"type:varchar(80);unique" form:"username" binding:"required"`
	Email          string `gorm:"type:varchar(80);unique" form:"email" binding:"required"`
	Password       []byte `gorm:"not null" json:"-"`
	IdRole    		uint  `gorm:"not null" json:"id_role"`
	Role Role `gorm:"foreignKey:RoleId"`
}

type Role struct {
	IdRole uint   `gorm:"primaryKey" json:"id_role"`
	Role   string `gorm:"type:varchar(80)"`
}

type AuthUserTokens struct {
	TokenId 	uint   `gorm:"primaryKey" json:"token_id"`
	AccessToken string `gorm:"unique" binding:"required"`
	RefeshToken string `gorm:"unique" binding:"required"`
	UserId uint  `gorm:"not null" json:"user_id"`
	Users Users `gorm:"foreignKey:TokenUserId"`
}

func (user *Users) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func (user *Users) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
