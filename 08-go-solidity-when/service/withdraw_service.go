package service

import "08-go-solidity-when/models"

type WithdrawService interface {
	GetOne(id int) (models.Withdraw, error)
	GetPage(withdraw models.Withdraw) ([]models.Withdraw, int64, int, int, error)
}

type withdrawService struct{}

func NewWithdrawService() WithdrawService {
	return withdrawService{}
}

func (withdrawService) GetOne(id int) (models.Withdraw, error) {
	withdraw := models.Withdraw{}
	err := models.DB.Where("id = ?", id).First(&withdraw).Error
	if err != nil {
		return withdraw, err
	}
	return withdraw, nil
}
func (withdrawService) GetPage(withdraw models.Withdraw) ([]models.Withdraw, int64, int, int, error) {
	tx := models.DB
	countTx := models.DB
	if withdraw.Id != 0 {
		tx = tx.Where("id = ?", withdraw.Id)
		countTx = countTx.Where("id = ?", withdraw.Id)
	}
	if withdraw.Src != "" {
		tx = tx.Where("dst = ?", withdraw.Src)
		countTx = countTx.Where("dst = ?", withdraw.Src)
	}
	if withdraw.Wad != "" {
		tx = tx.Where("wad = ?", withdraw.Wad)
		countTx = countTx.Where("wad = ?", withdraw.Wad)
	}
	if withdraw.Page.PageNumber == 0 {
		withdraw.Page.PageNumber = 1
	}
	if withdraw.Page.PageSize == 0 {
		withdraw.Page.PageSize = 10
	}
	var withdraws []models.Withdraw
	tx.Offset((withdraw.Page.PageNumber - 1) * withdraw.Page.PageSize).Limit(withdraw.Page.PageSize).Find(&withdraws)
	var count int64
	countTx.Table("withdraw").Count(&count)
	return withdraws, count, withdraw.Page.PageSize, withdraw.Page.PageNumber, nil
}
