package models

type Transfer struct {
	Id  int    `json:"id"`
	Src string `json:"src"`
	Dst string `json:"dst"`
	Wad string `gorm:"type:decimal(65,0)"`
}

func (Transfer) TableName() string {
	return "transfer"
}
