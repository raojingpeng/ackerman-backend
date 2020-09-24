package models

import (
	"backend/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Setup() {
	var err error
	databaseUrl := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		conf.Conf.Mysql.User,
		conf.Conf.Mysql.Password,
		conf.Conf.Mysql.Host,
		conf.Conf.Mysql.Port,
		conf.Conf.Mysql.Dbname)

	db, err = gorm.Open("mysql", databaseUrl)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Article{})
}
