package setting

/*
Package setting is used to parse the conf/app.ini and store the variable in the single struct
*/

import (
	"github.com/go-ini/ini"
	"path/filepath"

	//"Notify/utils/reflectu"
	"log"
	//"reflect"
	"time"
)

const (
	DebugMode = "deubg"
	ReleaseMode = "release"
	TestMode = "test"
)

type App struct {
	RuntimeRootPath string
	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting = &App{}


type Server struct{
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	JwtExpireTime time.Duration
	CacheSize  int
	CacheExpireTime time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	DbName string
	TablePrefix string
}
var DatabaseSetting = &Database{}

type Admin struct {
	Email string
	Password string
	UserId int
	Name string
}
var AdminSetting = &Admin{}



type Secret struct {
	JwtKey string
	JwtIssuer string
	SaltA string
	SaltB string
	AesKey string
	AesIv string
}
var SecretSetting = &Secret{}


// Setup init the setting struct, so before you use them, please
// use setting.Setup() to init them (only need once in the lifetime)
func Setup()  {
	Cfg,err:=ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatalf("Fail to parse `conf/app.ini` : %v",err)
	}

	//---------------- app config ----------------------
	err=Cfg.Section("app").MapTo(AppSetting)
	if err!=nil{
		log.Fatalf("Fail to parse 'AppSetting': %v", err)
	}
	// change the '/' to '\' in windows env, and do nothing in Unix
	AppSetting.RuntimeRootPath=filepath.FromSlash(AppSetting.RuntimeRootPath)
	AppSetting.LogSavePath=filepath.FromSlash(AppSetting.LogSavePath)


	//---------------- server config ----------------------
	err=Cfg.Section("server").MapTo(ServerSetting)
	if err!=nil	{
		log.Fatalf("Fail to parse 'ServerSetting': %v", err)
	}
	ServerSetting.ReadTimeout*=time.Second
	ServerSetting.WriteTimeout*=time.Second
	ServerSetting.JwtExpireTime*=time.Minute
	ServerSetting.CacheExpireTime*=time.Minute

	//---------------- database config ----------------------
	err=Cfg.Section("database").MapTo(DatabaseSetting)
	if err!=nil{
		log.Fatalf("Fail to parse 'DatabaseSetting': %v", err)
	}

	// you can use env which setting in docker(use os.env)

	//---------------- admin config ----------------------
	err=Cfg.Section("admin").MapTo(AdminSetting)
	if err!=nil{
		log.Fatalf("Fail to parse 'AdminSetting': %v", err)
	}


	err=Cfg.Section("secret").MapTo(SecretSetting)
	if err!=nil	{
		log.Fatalf("Fail to parse 'SecretSetting': %v", err)
	}

	/* 	you can use the following code to get env from docker

	reflectu.SetStructByReflect(reflect.ValueOf(&DatabaseSetting),"Type","Type")
	reflectu.SetStructByReflect(reflect.ValueOf(&SecretSetting),"JwtKey","JwtKey")

	*/

}