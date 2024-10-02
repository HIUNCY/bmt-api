package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"user_id"`
	FullName    string    `gorm:"type:varchar(100);not null" json:"full_name"`
	Email       string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNumber string    `gorm:"type:varchar(20);unique" json:"phone_number"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Accounts    []Account `gorm:"foreignKey:UserID" json:"accounts"`
}
