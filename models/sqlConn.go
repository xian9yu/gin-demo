package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"jwt/utils"

	"log"
)

var (
	DB *gorm.DB
	c  = *utils.NewCfg()
)

//初始化sql
func InitSQL() *gorm.DB {
	cfg := c.InitConfig()
	dsn := cfg.GetString("MySQL.DataSource")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "we_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `we_user`
		},
	})
	if err != nil {
		panic("MySQL启动异常")
	}

	if err = db.AutoMigrate(new(User)); err != nil {
		log.Println("同步数据库表失败:", err)

	}
	return db
}
