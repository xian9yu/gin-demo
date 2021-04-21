package models

type Article struct {
	Aid       int64  `json:"id" gorm:"primaryKey,comment:文章id"`
	Title     string `json:"title" gorm:"comment:文章标题"`
	CreatedAt int64  `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt int64  `gorm:"updated_at" gorm:"comment:编辑时间"`
	Content   string `json:"content" gorm:"comment:文章内容"`
	AuthorID  int64  `json:"authorId"  gorm:"comment:作者"`
	Category  string `json:"category" gorm:"comment:文章分类"`
	Status    string `json:"status"  gorm:"comment:是否可以查看 0隐藏 1显示"`
	IsComment int64  `json:"is_comment" gorm:"comment:是否可以评论 0关闭 1开启"`
	Views     int64  `json:"views" gorm:"comment:看过+1"`
}
