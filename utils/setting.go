package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	// 服务
	AppMode  string
	HttpPort string
	// 数据库
	DbUser        string
	DbPassWord    string
	DbHost        string
	DbPort        string
	DbName        string
	DbDefaultPage string
	DbDefaultSize string
	// STMP邮箱授权码
	EmailPassword string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err !=nil{
		log.Panicln("Error loading configuration file")
		return
	}
	LoadServer(file)
	LoadDb(file)
	LoadEmailStmp(file)
}

func LoadServer(file *ini.File){
	AppMode  =file.Section("server").Key("AppMode").MustString("debug")
	HttpPort =file.Section("server").Key("HttpPort").MustString(":9090")
}

func LoadDb(file *ini.File){
	DbUser        =file.Section("database").Key("DbUser").MustString("root")
	DbPassWord    =file.Section("database").Key("DbPassWord").MustString("123456")
	DbHost        =file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort        =file.Section("database").Key("DbPort").MustString("3306")
	DbName        =file.Section("database").Key("DbName").MustString("email_code")
	DbDefaultPage =file.Section("database").Key("DbDefaultPage").MustString("10")
	DbDefaultSize =file.Section("database").Key("DbDefaultSize").MustString("1")
}

func LoadEmailStmp(file *ini.File){
	EmailPassword = file.Section("email").Key("EmailPassword").MustString("") // STMP授权码
}