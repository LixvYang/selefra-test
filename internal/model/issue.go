package model

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Issue struct {
	// uid == github_id
	Uid          string `gorm:"type:varchar(255);not null" json:"uid"`
	User         User   `gorm:"foreignKey:uid" json:"user"`
	IssureNumber string
	Body         string
	TokenNum     decimal.Decimal
}

func (*Issue) CreateIssue(data *Issue) (err error) {
	if err = db.Create(&data).Error; err != nil {
		return errors.New("CreateIssue error: " + err.Error())
	}
	return nil
}

func (*Issue) GetIssue(data *Issue) (issue Issue, err error) {
	if err = db.Select("uid, number, body, ").Find(&issue).Error; err != nil {
		return issue, errors.New("GetIssue error: " + err.Error())
	}
	return issue, nil
}
