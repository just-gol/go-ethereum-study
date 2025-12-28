package models

type Withdraw struct {
	Id  int    `json:"id"`
	Src string `json:"src"`
	Wad string `gorm:"type:decimal(65,0)"`
}

func (Withdraw) TableName() string {
	return "withdraw"
}
