package service

import "08-go-solidity-when/models"

type TransferService interface {
	GetOne(id int) (models.Transfer, error)
	GetPage(transfer models.Transfer) ([]models.Transfer, int64, int, int, error)
}

type transferService struct{}

func NewTransferService() TransferService {
	return transferService{}
}

func (transferService) GetOne(id int) (models.Transfer, error) {
	transfer := models.Transfer{}
	err := models.DB.Where("id = ?", id).First(&transfer).Error
	if err != nil {
		return transfer, err
	}
	return transfer, nil
}
func (transferService) GetPage(transfer models.Transfer) ([]models.Transfer, int64, int, int, error) {
	tx := models.DB
	countTx := models.DB
	if transfer.Id != 0 {
		tx = tx.Where("id = ?", transfer.Id)
		countTx = countTx.Where("id = ?", transfer.Id)
	}
	if transfer.Dst != "" {
		tx = tx.Where("dst = ?", transfer.Dst)
		countTx = countTx.Where("dst = ?", transfer.Dst)
	}
	if transfer.Dst != "" {
		tx = tx.Where("wad = ?", transfer.Dst)
		countTx = countTx.Where("wad = ?", transfer.Dst)
	}
	if transfer.Wad != "" {
		tx = tx.Where("wad = ?", transfer.Wad)
		countTx = countTx.Where("wad = ?", transfer.Wad)
	}
	if transfer.Page.PageNumber == 0 {
		transfer.Page.PageNumber = 1
	}
	if transfer.Page.PageSize == 0 {
		transfer.Page.PageSize = 10
	}
	var transfers []models.Transfer
	tx.Offset((transfer.Page.PageNumber - 1) * transfer.Page.PageSize).Limit(transfer.Page.PageSize).Find(&transfers)
	var count int64
	countTx.Table("transfer").Count(&count)
	return transfers, count, transfer.Page.PageSize, transfer.Page.PageNumber, nil
}
