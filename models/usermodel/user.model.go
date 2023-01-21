package usermodel

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Username  string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(255)"`
	Role      string `gorm:"type:varchar(20)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Register struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
