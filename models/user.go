package models

import "time"

type UserModel struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"varchar(100)" json:"first_name"`
	LastName  string    `gorm:"varchar(100)" json:"last_name"`
	Address   string    `gorm:"text" json:"address"`
	Phone     string    `gorm:"varchar(20)" json:"phone"`
	CreatedAt time.Time `gorm:"timestamp; default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt time.Time `gorm:"timestamp; default:CURRENT_TIMESTAMP()" json:"updated_at"`
}