package models

import "time"

type Transfer struct {
	ID            uint      `gorm:"primaryKey" json:"transfer_id"`
	FromAccountID uint      `gorm:"not null" json:"sender_id"`
	ToAccountID   uint      `gorm:"not null" json:"receiver_id"`
	Amount        float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	TransferDate  time.Time `gorm:"autoCreateTime"`
	Description   string    `gorm:"type:text" json:"description"`
}
