package models

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey" json:"transaction_id"`
	AccountID       uint      `gorm:"not null" json:"account_id"`
	TransactionType string    `gorm:"type:enum('deposit','withdrawal','transfer');not null" json:"transaction_type"`
	Amount          float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	TransactionDate time.Time `gorm:"autoCreateTime"`
	Description     string    `gorm:"type:text" json:"description"`
}
