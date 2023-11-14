package models

import "time"

type AuthModel struct {
	Id       	int64  `gorm:"primaryKey" json:"id"`
	Username 	string `gorm:"varchar(30)" json:"username"`
	Email    	string `gorm:"varchar(50)" json:"email"`
	Password 	string `gorm:"carchar(30)" json:"password"`
	RoleId   	int64  `gorm:"int" json:"role_id"`
	Role 		Role `gorm:"foreignKey:RoleId"`
	CreatedAt 	time.Time	`gorm:"timestamp; default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt 	time.Time	`gorm:"timestamp; default:CURRENT_TIMESTAMP()" json:"updated_at"`
}