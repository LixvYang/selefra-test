package model

import "gorm.io/gorm"

type Email struct {
	GithubID  string
	Name      string
	EmailLink string
	gorm.Model
}

func (*Email) CreateEmail() {
}
