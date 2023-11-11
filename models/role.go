package models

import "time"

type RoleModel struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	RoleName  string `gorm:"varchar(50)" json:"role_name"`
	CreatedAt time.Time	`gorm:"timestamp; default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt time.Time	`gorm:"timestamp; default:CURRENT_TIMESTAMP()" json:"updated_at"`
}