package models

import "time"

type User struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"varchar(100)" json:"first_name"`
	LastName  string    `gorm:"varchar(100)" json:"last_name"`
	Address   string    `gorm:"text" json:"address"`
	Phone     string    `gorm:"varchar(20)" json:"phone"`
	CreatedAt time.Time `gorm:"timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"timestamp" json:"updated_at"`
}