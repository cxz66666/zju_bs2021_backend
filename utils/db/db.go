package db

import (
	"annotation/utils/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	MysqlDB  *gorm.DB
)

var (
	//  MySqlDNS = `user:123456@tcp(10.79.25.200:3306)/notify?charset=utf8mb4&parseTime=True&loc=Local`
	MySqlDNS  = `%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`

	// MySqlDNS = `root:123@tcp(127.0.0.1:3306)/notify?charset=utf8mb4&parseTime=True&loc=Local`
)

func Setup() {
	//连接mysql数据库
	var (
		dbType, dbName, user, password, host, tablePrefix string
		err error
	)
	dbType=setting.DatabaseSetting.Type
	dbName=setting.DatabaseSetting.DbName
	user=setting.DatabaseSetting.User
	password=setting.DatabaseSetting.Password
	host=setting.DatabaseSetting.Host
	tablePrefix=setting.DatabaseSetting.TablePrefix

	// debug mode

	MySqlDNS = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password,host,dbName)

	MysqlDB, err = gorm.Open(mysql.Open(MySqlDNS), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to init db")
	}


	if len(tablePrefix)>0{
		fmt.Printf("[warning] tablePrefix '%s' will be nothing to do in current version",tablePrefix)
	}
	if dbType!="mysql"{
		fmt.Printf("[warning]  '%s' will be not be use in current version",dbType)
		os.Exit(-1)
	}
	// auto migrate  it can't handle the dependency relations, so you need handle it by yourself
	// MysqlDB.AutoMigrate(&user.User{})
	// MysqlDB.AutoMigrate(&subscription.Subscription{})
	// MysqlDB.AutoMigrate(&subscription.Order{})
	// MysqlDB.AutoMigrate(&notice.Notice{})
}

