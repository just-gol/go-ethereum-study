package models

type SyncState struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:64;uniqueIndex"`
	BlockNumber uint64 `gorm:"not null"`
}

func (SyncState) TableName() string {
	return "sync_state"
}
