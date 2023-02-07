package email

import "gorm.io/gorm"

type Email struct {
	ToEmailLink string
	gorm.Model
}
