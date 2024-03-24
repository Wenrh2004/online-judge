package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string `gorm:"column:id;type:varchar(36);primarykey"`
	Username  string `gorm:"column:username;type:varchar(20);not null;unique"`
	Password  string `gorm:"column:password;type:varchar(255);not null;unique"`
	Role      int32  `gorm:"column:role;type:tinyint;not null;default:0"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);"`
	Phone     string `gorm:"column:phone;type:varchar(11);"`
	Email     string `gorm:"column:email;type:varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
