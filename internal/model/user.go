package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	GithubID   string
	GithubName string
	PublicKey  string
	AvatarUrl  string
	Email      string
	gorm.Model
}

// 增删查改
func (*User) CreateUser(data *User) (err error) {
	if err = db.Create(&data).Error; err != nil {
		return errors.New("CreateUser error: " + err.Error())
	}
	return nil
}

// Delete user
func (*User) DeleteUser(data *User) (err error) {
	if err = db.Where("github_id = ?", data.GithubID).Delete(User{}).Error; err != nil {
		return errors.New("DeleteUser error: " + err.Error())
	}
	return nil
}

// CheckUser 查询用户是否存在
func (*User) CheckUser(data *User) (err error) {
	var user User
	err = db.Model(&User{}).Where("github_id = ?", data.GithubID).Last(&user).Error
	if user.GithubID == "" {
		return errors.New("CheckUser error: " + err.Error())
	}
	return nil
}

func (*User) UpdateUser(data *User) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return errors.New("UpdateUser error: " + err.Error())
	}

	// 锁住指定 github_id 的 User 记录
	if err = tx.Set("gorm:query_option", "FOR UPDATE").Where("github_id = ?", data.GithubID).Error; err != nil {
		tx.Rollback()
		return errors.New("UpdateUser error: " + err.Error())

	}

	var maps = make(map[string]interface{})
	maps["avatar_url"] = data.AvatarUrl
	maps["email"] = data.Email
	maps["github_name"] = data.GithubName
	maps["public_key"] = data.PublicKey
	if err = db.Model(&User{}).Where("github_id = ? ", data.GithubID).Updates(maps).Error; err != nil {
		return errors.New("UpdateUser error: " + err.Error())

	}
	if err = tx.Commit().Error; err != nil {
		return errors.New("UpdateUser error: " + err.Error())

	}
	return nil
}
