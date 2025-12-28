package models

type Approval struct {
	Id  int    `json:"id"`
	Src string `json:"src"`
	Guy string `json:"guy"`
	Wad string `gorm:"type:decimal(65,0)"`
}

func (Approval) TableName() string {
	return "approval"
}
