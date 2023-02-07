package utils

import (
	"log"
	"os"
	"sync"

	"gopkg.in/ini.v1"
)

var (
	DomainName string
	IP         string
	AppMode    string
	HttpPort   string
	JwtKey     string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	MailHost string
	MailPort string
	MailUser string
	MailPass string
)

var InitSetting sync.Once

func Init() {
	str, _ := os.Getwd()
	f, err := ini.Load(str + "/internal/config/config.ini")
	if err != nil {
		log.Printf("配置文件读取错误:%s", err)
	}
	LoadServer(f)
	LoadData(f)
	LoadMail(f)
}

func LoadServer(file *ini.File) {
	DomainName = file.Section("server").Key("DomainName").MustString("")
	IP = file.Section("server").Key("IP").MustString("")
	AppMode = file.Section("server").Key("AppMode").MustString("")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("")
	DbHost = file.Section("database").Key("DbHost").MustString("")
	DbPort = file.Section("database").Key("DbPort").MustString("")
	DbUser = file.Section("database").Key("DbUser").MustString("")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("")
}

func LoadMail(file *ini.File) {
	MailHost = file.Section("mail").Key("MailHost").MustString("")
	MailPort = file.Section("mail").Key("MailHost").MustString("")
	MailUser = file.Section("mail").Key("MailUser").MustString("")
	MailPass = file.Section("mail").Key("MailPass").MustString("")
}
