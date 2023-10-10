package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	DbConfig DbConfig
}

type DbConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Name         string
	Charset      string
	Driver       string
	MaxIdleConns int
	MaxOpenConns int
}

var (
	Cg           Config
	ErrorFileMsg string = "配置文件读取错误，请检查文件路径:"
	FileDir      string
)

// ReadConfig 从config.ini中读取配置
func ReadConfig() {
	file, err := ini.Load(FileDir)
	if err != nil {
		fmt.Println(ErrorFileMsg, err)
	}
	LoadMysql(file)
}

// LoadMysql 从config.ini中读取mysql配置
func LoadMysql(file *ini.File) {
	section := file.Section("mysql")
	Cg.DbConfig.Host = section.Key("DbHost").String()
	Cg.DbConfig.Port, _ = section.Key("DbPort").Int()
	Cg.DbConfig.User = section.Key("DbUser").String()
	Cg.DbConfig.Password = section.Key("DbPassword").String()
	Cg.DbConfig.Name = section.Key("DbName").String()
	Cg.DbConfig.Charset = section.Key("Charset").String()
	Cg.DbConfig.Driver = section.Key("Driver").String()
	Cg.DbConfig.MaxIdleConns, _ = section.Key("MaxIdleConns").Int()
	Cg.DbConfig.MaxOpenConns, _ = section.Key("MaxOpenConns").Int()
}

// init 读取配置文件
func init() {
	dir, _ := os.Getwd()
	// 判断是否是测试环境
	if dir[len(dir)-5:] == "tests" {
		// 返回上一级目录
		dir = dir[:len(dir)-5]
	} else {
		// 拼接 /
		dir = dir + "/"
	}
	FileDir = dir + "config/config.ini"
	ReadConfig()
}
