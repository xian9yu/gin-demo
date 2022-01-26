package models

const PasswordMd5Key = "passwordMd5SecureKey"

type User struct {
	Uid         uint64 `json:"uid" gorm:"size:12;primaryKey;unique;notnull;comment:用户id"`
	Username    string `json:"username" gorm:"size:60;comment:用户名"`
	Mail        string `json:"mail" gorm:"size:60;notnull;comment:邮箱"`
	Password    string `json:"-"  gorm:"size:33;notnull;comment:用户登录密码"`
	Url         string `json:"url" gorm:"size:60;comment:网站url"`
	Group       string `json:"group" gorm:"size:15;notnull;comment:用户分组(管理员,用户)"`
	Status      string `json:"status" gorm:"size:6;notnull;comment:账户状态(启用,关闭)"`
	CreatedTime uint64 `json:"created_time" gorm:"autoCreateTime;notnull comment:user创建时间"`
	UpdatedTime uint64 `json:"updated_time" gorm:"autoUpdateTime;comment:上一次修改信息时间"`
	LastLogin   uint64 `json:"last_login" gorm:"comment:上次登录时间"`
}

// Login 用户登录
func (u *User) Login() bool {
	var count int64
	DB.Model(&User{}).Where("mail = ? AND password = ?", u.Mail, u.Password).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// Register 注册
func (u *User) Register(user User) (int64, uint64, error) {
	result := DB.Create(&user)
	return result.RowsAffected, user.Uid, result.Error
}

//func (u *User) FindById(id int) (users *User, err error) {
//	var user User
//	err = DB.Where("`id` = ?", id).Find(&user).Error
//	return &user, err
//}
//
//func (u *User) FindByName(username string) (users *User, err error) {
//	var user User
//	err = DB.Where("`user_name` = ?", username).First(&user).Error
//	return &user, err
//}

func (u *User) GetAll() (total int64, list []User, err error) {
	var userList []User

	err = DB.Model(&u).Count(&total).Error
	err = DB.Find(&userList).Error
	return total, userList, err
}

//  更新
func (u *User) Update(user User) {
	DB.Model(&user).Updates(user)

}

type getUserInfoGenerics interface {
	int | string
}

func GetInfoBy[T getUserInfoGenerics](query string, v T) (users *User, err error) {
	var user User
	err = DB.Where("`"+query+"` = ?", v).Find(&user).Error
	return &user, err
}
