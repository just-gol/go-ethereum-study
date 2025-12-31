package models

type Withdraw struct {
	Id   int    `json:"id"`
	Src  string `json:"src"`
	Wad  string `gorm:"type:decimal(65,0)"`
	Page Page   `json:"-" gorm:"-"`
}

func (Withdraw) TableName() string {
	return "withdraw"
}
