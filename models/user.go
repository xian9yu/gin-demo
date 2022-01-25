package models

type User struct {
	Uid         uint64 `json:"uid" gorm:"size:12;primaryKey;unique;notnull;comment:用户id"`
	Username    string `json:"username" gorm:"size:60;comment:用户名"`
	Mail        string `json:"mail" gorm:"size:60;notnull;comment:邮箱"`
	Password    string `json:"-"  gorm:"size:33;notnull;comment:用户登录密码"`
	Url         string `json:"url" gorm:"size:60;comment:网站url"`
	Group       string `json:"group" gorm:"size:33;notnull;comment:用户分组"`
	CreatedTime uint64 `json:"created_time" gorm:"autoCreateTime;notnull comment:user创建时间"`
	UpdatedTime uint64 `json:"updated_time" gorm:"autoUpdateTime;comment:上一次修改信息时间"`
	LastLogin   uint64 `json:"last_login" gorm:"comment:上一次登录时间"`
}

// Login 用户登录
func (u *User) Login() (err error) {
	var user User
	return DB.Where("user_name = ? AND pass_word = ?", u.Username, u.Password).Take(&user).Error
}

// Register 用户注册
func (u *User) Register(user *User) (int64, uint64, error) {
	result := DB.Create(&user)
	return result.RowsAffected, user.Uid, result.Error
}

func (u *User) FindById(id int) (users *User, err error) {
	var user User
	err = DB.Where("`id` = ?", id).Find(&user).Error
	return &user, err
}

func (u *User) FindByName(username string) (users *User, err error) {
	var user User
	err = DB.Where("`user_name` = ?", username).First(&user).Error
	return &user, err
}

func (u *User) GetAll() (total int64, list []User, err error) {
	var userList []User

	err = DB.Model(&u).Count(&total).Error
	err = DB.Find(&userList).Error
	return total, userList, err
}

type info interface {
	int | string
}

func GetInfoBy[T info](query string, v T) (users *User, err error) {
	var user User
	err = DB.Where("`"+query+"` = ?", v).Find(&user).Error
	return &user, err
}
