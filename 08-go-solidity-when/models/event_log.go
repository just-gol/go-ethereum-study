package models

import "time"

type EventLog struct {
	ID          uint      `gorm:"primaryKey"`
	TxHash      string    `gorm:"size:66;not null;uniqueIndex:uniq_tx_log"`
	LogIndex    uint      `gorm:"not null;uniqueIndex:uniq_tx_log"`
	BlockNumber uint64    `gorm:"not null"`
	Event       string    `gorm:"size:32;not null"`
	Contract    string    `gorm:"size:42;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (EventLog) TableName() string {
	return "event_log"
}
