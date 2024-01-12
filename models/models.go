package models

import (
	"golang.org/x/crypto/bcrypt"
)
type Users struct {
	Id 			uint  `gorm:"primaryKey" json:"id"`
	Username       string `gorm:"type:varchar(80);unique" form:"username" binding:"required"`
	Email          string `gorm:"type:varchar(80);unique" form:"email" binding:"required"`
	Password       []byte `gorm:"not null" json:"-"`
	IdRole    		uint  `gorm:"not null" json:"roleId"`
	Role Role `gorm:"foreignKey:IdRole"`
}

type Role struct {
	Id uint   `gorm:"primaryKey" json:"id"`
	Role   string `gorm:"type:varchar(80)"`
}

type AuthUserTokens struct {
	Id 	uint   `gorm:"primaryKey" json:"id"`
	AccessToken string `gorm:"type:varchar(350);" form:"access_token" binding:"required"`
	RefeshToken string `gorm:"type:varchar(350);" form:"refesh_token" binding:"required"`
	UserId uint  `gorm:"not null" json:"usersId"`
	Users Users `gorm:"foreignKey:UserId"`
}

type Coffe struct {
	Id uint `gorm:"primaryKey" json:"id"`
	JenisCoffe string `gorm:"type:varchar(350);" form:"jenis_coffe" binding:"required"`
	HargaCoffe string `gorm:"type:varchar(350);" form:"harga_coffe" binding:"required"`
	ImagesCoffe string `gorm:"type:varchar(350);" form:"images_coffe" binding:"required"`
	UrlImageCoffe string `gorm:"type:varchar(350);" form:"url_image_coffe" binding:"required"`
}

func (users *Users) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	users.Password = hashedPassword
	return nil
}

func (users *Users) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(users.Password, []byte(password))
}