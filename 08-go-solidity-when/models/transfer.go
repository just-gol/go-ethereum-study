package models

type Transfer struct {
	Id   int    `json:"id"`
	Src  string `json:"src"`
	Dst  string `json:"dst"`
	Wad  string `gorm:"type:decimal(65,0)"`
	Page Page   `json:"-" gorm:"-"`
}

func (Transfer) TableName() string {
	return "transfer"
}
