package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//TODO :viper bug：添加多路径配置文件的时候不能通过结构体读取config文件

func NewCfg() *Config {
	return &Config{}
}

type Config struct {
	ListenOn struct {
		Host string
		Port string
	}
	SQL struct {
		Driver     string
		DataSource string
	}
	JWT struct {
		SecretKey string
		ExpireAt  int64
	}
}

//var (
//	readError      = errors.New("读取配置文件失败")
//	unMarshalError = errors.New("解析配置文件失败")
//)

// 获取相对路径
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
var (
	file, _ = exec.LookPath(os.Args[0])
	path, _ = filepath.Abs(file)
	index   = strings.LastIndex(path, string(os.PathSeparator))

	c1 = path[:index]          //相对路径
	c2 = getCurrentDirectory() //相对路径
	c3 = "."                   //绝对路径
)

func (c *Config) InitConfig() *viper.Viper {
	v := viper.New()

	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//index := strings.LastIndex(path, string(os.PathSeparator))
	//
	//c1 := path[:index]          //相对路径
	//c2 := getCurrentDirectory() //相对路径
	//c3 := "."                   //绝对路径

	//添加配置文件可能存在的路径
	v.AddConfigPath(c1)
	v.AddConfigPath(c2)
	v.AddConfigPath(c3)

	//配置文件名称
	v.SetConfigName("vconf")
	//配置文件类型
	v.SetConfigType("yaml")

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败:", err)
	}

	// 将文件内容解析后封装到cfg对象中
	err = v.Unmarshal(&Config{})
	if err != nil {
		fmt.Println("解析配置文件失败:", err)
	}
	return v
}
