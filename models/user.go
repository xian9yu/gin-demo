package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey,comment:用户id"`
	Username  string `json:"username" gorm:"comment:用户登录名"`
	Password  string `json:"-"  gorm:"comment:用户登录密码"`
	Mail      string `json:"mail" gorm:"comment:邮箱"`
	Url       string `json:"url" gorm:"comment:网站url"`
	Nickname  string `json:"nickname" gorm:"comment:用户显示名称"`
	Group     string `json:"group" gorm:"comment:用户分组"`
	CreatedAt int64  `gorm:"created_at" gorm:"comment:创建时间"`
	UpdatedAt int64  `gorm:"updated_at" gorm:"comment:上一次修改信息时间"`
	Logged    int64  `json:"logged" gorm:"comment:上一次登录时间"`
}

// Login 用户登录
func (u *User) Login() (err error) {
	var user User
	return DB.Where("user_name = ? AND pass_word = ?", u.Username, u.Password).Take(&user).Error
}

// Register 用户注册
func (u *User) Register(username string) (err error) {
	//如果没有查找到记录则根据结构体创建
	err = DB.Find(&u, "user_name = ?", username).Create(&u).Error
	return err
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

func (u *User) GetList() (total int64, list []User, err error) {
	var userList []User

	err = DB.Model(&u).Count(&total).Error
	err = DB.Find(&userList).Error
	return total, userList, err
}
