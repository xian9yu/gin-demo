package models

import (
	"errors"
)

type User struct {
	ID        int64  `json:"id" gorm:"comment:用户ID"`
	UserName  string `json:"user_name" gorm:"comment:用户登录名"`
	PassWord  string `json:"-"  gorm:"comment:用户登录密码"`
	CreatedAt int64  `gorm:"created_at"`
	UpdatedAt int64  `gorm:"updated_at"`
}

// 用户登录
func (u *User) Login() (err error) {
	var user User
	return DB.Where("user_name = ? AND pass_word = ?", u.UserName, u.PassWord).Take(&user).Error
}

//用户注册
func (u *User) Register(username string) error {
	//如果没有查找到记录则根据结构体创建
	err := DB.Find(&u, "user_name = ?", username).FirstOrCreate(&u).Error
	if err != nil { // 判断用户名是否注册
		return errors.New("用户名已注册")
	}
	return nil
}

func (u *User) FindById(id int) (users *User, err error) {
	var user User
	err = DB.Where("`id` = ?", id).First(&user).Error
	return &user, err
}

func (u *User) FindByName(username string) (users *User, err error) {
	var user User
	err = DB.Where("`user_name` = ?", username).First(&user).Error
	return &user, err
}

//get List
func (u *User) GetList() (total int64, list []User, err error) {
	var userList []User

	err = DB.Model(&u).Count(&total).Error
	err = DB.Find(&userList).Error
	return total, userList, err
}
