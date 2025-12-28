package models

type Deposit struct {
	Id  int    `json:"id"`
	Dst string `json:"dst"`
	Wad string `gorm:"type:decimal(65,0)"`
}

func (Deposit) TableName() string {
	return "deposit"
}
