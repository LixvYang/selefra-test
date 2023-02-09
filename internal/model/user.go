/*
 * @description:
 * @param:
 * @return:
 */
package model

import (
	"errors"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	GithubID   string
	GithubName string
	PublicKey  string
	AvatarUrl  string
	EmailLink  string
	TokenNum   decimal.Decimal `gorm:"type:decimal(36,18);default 0" json:"token_num"`
	// 若没有绑定PublicKey, 临时存储token
	TempToken decimal.Decimal `gorm:"type:decimal(36,18);default 0" json:"temp_token"`
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
	maps["email"] = data.EmailLink
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

func (*User) IncrUserToken(data *User, num decimal.Decimal) (err error) {
	if err = db.Where("github_id = ?", data.GithubID).Exec("set token = token + ?", num).Error; err != nil {
		return errors.New("IncrUserToken error: " + err.Error())
	}
	return nil
}

// 查询某个人的token数目
func (*User) GetUserTokenNum(data *User) (num decimal.Decimal, err error) {
	var user User
	if err = db.Where("github_id = ?", data.GithubID).Find(&user).Error; err != nil {
		return decimal.New(0, 0), errors.New("GetUserTokenNum error: " + err.Error())
	}
	return user.TokenNum, nil
}

func (*User) ListUsers() ([]User, error) {
	// List users
	var users []User
	if err := db.Model(&User{}).Error; err != nil {
		return users, err
	}

	if err := db.Where("").Order("id desc").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (*User) DescUserTokenNum(data *User, num decimal.Decimal) (err error) {
	if err = db.Where("github_id = ?", data.GithubID).Exec("set token_num = token_num - ?", num).Error; err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (*User) CheckUserBind(data *User) bool {
	var u User
	var err error
	if err = db.Where("github_id = ?", data.GithubID).First(&u).Error; err != nil {
		return false
	}
	if u.PublicKey == "" {
		return false
	}
	return true
}

// 增加临时存储token
func (*User) IncrUserTempToken(data *User, num decimal.Decimal) (err error) {
	if err = db.Where("github_id = ?", data.GithubID).Exec("set temp_token = temp_token + ?", num).Error; err != nil {
		return errors.New("IncrUserTempToken" + err.Error())
	}
	return nil
}
