package Databases

import (
	"backend/Config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Mysql *gorm.DB

func init() {
	var err error
	databaseUrl := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Config.Conf.Mysql.User,
		Config.Conf.Mysql.Password,
		Config.Conf.Mysql.Host,
		Config.Conf.Mysql.Port,
		Config.Conf.Mysql.Dbname)

	Mysql, err = gorm.Open("mysql", databaseUrl)
	if err != nil {
		panic("Failed to connect database")
	}
	if Mysql.Error != nil {
		panic(fmt.Errorf("Database error: %s \n", Mysql.Error))
	}
}
