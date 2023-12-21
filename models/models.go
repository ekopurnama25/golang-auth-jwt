package models

type Users struct {
	UserId 			uint   `gorm:"primaryKey" json:"user_id"`
	Username       string `gorm:"type:varchar(80);unique" form:"username" binding:"required"`
	Email          string `gorm:"type:varchar(80);unique" form:"email" binding:"required"`
	Password       []byte `gorm:"not null" json:"-"`
	RoleId         string `gorm:"type:varchar(100)"`
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
	TokenUserId string `gorm:"type:varchar(100)"`
	Users Users `gorm:"foreignKey:TokenUserId"`
}