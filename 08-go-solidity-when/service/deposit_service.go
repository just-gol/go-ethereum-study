package service

import "08-go-solidity-when/models"

type DepositService interface {
	GetOne(id int) (models.Deposit, error)
	GetPage(deposit models.Deposit) ([]models.Deposit, int64, int, int, error)
}

type depositService struct{}

func NewDepositService() DepositService {
	return depositService{}
}

func (depositService) GetOne(id int) (models.Deposit, error) {
	deposit := models.Deposit{}
	err := models.DB.Where("id = ?", id).First(&deposit).Error
	if err != nil {
		return deposit, err
	}
	return deposit, nil
}
func (depositService) GetPage(deposit models.Deposit) ([]models.Deposit, int64, int, int, error) {
	tx := models.DB
	countTx := models.DB
	if deposit.Id != 0 {
		tx = tx.Where("id = ?", deposit.Id)
		countTx = countTx.Where("id = ?", deposit.Id)
	}
	if deposit.Dst != "" {
		tx = tx.Where("dst = ?", deposit.Dst)
		countTx = countTx.Where("dst = ?", deposit.Dst)
	}
	if deposit.Wad != "" {
		tx = tx.Where("wad = ?", deposit.Wad)
		countTx = countTx.Where("wad = ?", deposit.Wad)
	}
	if deposit.Page.PageNumber == 0 {
		deposit.Page.PageNumber = 1
	}
	if deposit.Page.PageSize == 0 {
		deposit.Page.PageSize = 10
	}
	var deposits []models.Deposit
	tx.Offset((deposit.Page.PageNumber - 1) * deposit.Page.PageSize).Limit(deposit.Page.PageSize).Find(&deposits)
	var count int64
	countTx.Table("deposit").Count(&count)
	return deposits, count, deposit.Page.PageSize, deposit.Page.PageNumber, nil
}
