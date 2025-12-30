package service

import "08-go-solidity-when/models"

type ApprovalService interface {
	GetOne(id int) (models.Approval, error)
	GetPage(approval models.Approval) ([]models.Approval, int64, int, int, error)
}

type approvalService struct{}

func NewApprovalService() ApprovalService {
	return &approvalService{}
}

func (s *approvalService) GetOne(id int) (models.Approval, error) {
	approval := models.Approval{}
	err := models.DB.Where("id = ?", id).First(&approval).Error
	if err != nil {
		return approval, err
	}
	return approval, nil
}

func (s *approvalService) GetPage(approval models.Approval) ([]models.Approval, int64, int, int, error) {
	pageNumber := approval.Page.PageNumber
	if pageNumber == 0 {
		pageNumber = 1
	}
	pageSize := approval.Page.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	var approvalList []models.Approval
	tx := models.DB
	txCount := models.DB.Table("approval")
	if approval.Id != 0 {
		tx = tx.Where("id = ?", approval.Id)
		txCount = txCount.Where("id = ?", approval.Id)
	}
	if approval.Src != "" {
		tx = tx.Where("src = ?", approval.Src)
		txCount = txCount.Where("src = ?", approval.Src)
	}
	if approval.Guy != "" {
		tx = tx.Where("guy = ?", approval.Guy)
		txCount = txCount.Where("guy = ?", approval.Guy)
	}
	if approval.Wad != "" {
		tx = tx.Where("wad = ?", approval.Wad)
		txCount = txCount.Where("wad = ?", approval.Wad)
	}
	tx.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&approvalList)
	var count int64
	txCount.Count(&count)
	return approvalList, count, pageSize, pageNumber, nil
}
