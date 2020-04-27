package Databases

import (
	"backend/Config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	databaseUrl := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Config.Conf.Db.User, Config.Conf.Db.Password, Config.Conf.Db.Host, Config.Conf.Db.Port, Config.Conf.Db.Dbname)

	DB, err = gorm.Open("mysql", databaseUrl)
	if err != nil {
		panic("Failed to connect database")
	}
	if DB.Error != nil {
		panic(fmt.Errorf("Database error: %s \n", DB.Error))
	}
}
