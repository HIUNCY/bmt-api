package models

import "time"

type Account struct {
	ID            uint          `gorm:"primaryKey" json:"account_id"`
	UserID        uint          `gorm:"not null" json:"user_id"`
	AccountNumber string        `gorm:"type:varchar(20);unique;not null" json:"account_number"`
	Balance       float64       `gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime"`
	Transactions  []Transaction `gorm:"foreignKey:AccountID" json:"transactions"`
	FromTransfers []Transfer    `gorm:"foreignKey:FromAccountID" json:"from_transfers"`
	ToTransfers   []Transfer    `gorm:"foreignKey:ToAccountID" json:"to_transfers"`
}
