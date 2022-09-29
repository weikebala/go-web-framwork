package main

import (
	"fmt"
	"mygin/orm"
)
func insertOne(name string, age uint8){
	user := orm.User{Name: name, Age: 18}

	result := orm.DB.Create(&user) // 通过数据的指针来创建
	fmt.Println(user, result.RowsAffected)
}
func insertMuti(){
	var users = []orm.User{{Name: "jinzhu2"}, {Name: "jinzhu3"}, {Name: "jinzhu4"}}
	orm.DB.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID) // 1,2,3
	}
}

func main() {
	//insertOne("a", 17)
	insertMuti()
}
