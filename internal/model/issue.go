package model

import (
	"errors"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Issue struct {
	// uid == github_id
	gorm.Model
	Uid         string `gorm:"type:varchar(255);not null" json:"uid"`
	User        User   `gorm:"foreignKey:uid" json:"user"`
	IssueNumber string
	Body        string
	TokenNum    decimal.Decimal
}

func (*Issue) CreateIssue(data *Issue) (err error) {
	if err = db.Create(&data).Error; err != nil {
		return errors.New("CreateIssue error: " + err.Error())
	}
	return nil
}

func (*Issue) GetIssue(data *Issue) (Issue, error) {
	var issue Issue
	if err := db.Select("uid, body, token_num").Where("issue_number = ?", data.IssueNumber).Find(&issue).Error; err != nil {
		return issue, errors.New("GetIssue error: " + err.Error())
	}
	return issue, nil
}
