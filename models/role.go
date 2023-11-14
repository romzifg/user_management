package models

import "time"

type Role struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	RoleName  string `gorm:"varchar(50)" json:"role_name"`
	CreatedAt time.Time	`gorm:"timestamp" json:"created_at"`
	UpdatedAt time.Time	`gorm:"timestamp" json:"updated_at"`
}