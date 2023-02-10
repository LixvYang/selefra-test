package model

import "gorm.io/gorm"

type Email struct {
	gorm.Model

	GithubID  string
	Name      string
	EmailLink string
}

func (*Email) CreateEmail() {
}
