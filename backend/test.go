package main

import (
	"backend/Databases"
	"backend/Models"
	"fmt"
)

func main() {
	//user:= Models.User{
	//	Username:        "raojingpeng",
	//	Email:           "withrjp@gmail.com",
	//}

	//Databases.DB.Create(&user)
	//fmt.Print("ok")
	user2 := Models.User{}
	Databases.DB.Where("username = ?", "raojingpeng").First(&user2)
	fmt.Println(user2.CheckPassword("rjp1994"))
	fmt.Println(user2.CheckPassword("rjo1994."))
	fmt.Println(user2.CheckPassword("rjp1994."))
}
