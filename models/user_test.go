package models

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	user := new(User)
	rows, uid, err := user.Register(User{
		Mail:     "xx@qq.com",
		Password: "123456",
		Group:    "普通",
		Status:   "启用",
	})
	fmt.Println(user)
	if err != nil {
		fmt.Println("err= ", err)
	}
	fmt.Println("rows= ", rows)
	fmt.Println("uid= ", uid)
}

func TestUpdate(t *testing.T) {
	user := new(User)
	user.Update(User{
		Uid:         0,
		Username:    "",
		Mail:        "",
		Password:    "",
		Url:         "",
		Group:       "",
		Status:      "",
		CreatedTime: 0,
		UpdatedTime: 0,
		LastLogin:   0,
	})
}
