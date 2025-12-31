package models

type Deposit struct {
	Id   int    `json:"id"`
	Dst  string `json:"dst"`
	Wad  string `gorm:"type:decimal(65,0)"`
	Page Page   `json:"-" gorm:"-"`
}

func (Deposit) TableName() string {
	return "deposit"
}
