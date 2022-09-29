package main

import (
	"encoding/json"
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

var user orm.User

func first() []byte {
	orm.DB.First(&user)
	json,_ := json.Marshal(user)
	return json
}

func findAll(){
	user_sli := make([]orm.User, 1)
	orm.DB.Find(&user_sli)
	json,_ := json.Marshal(user_sli)
	fmt.Println(string(json))
}

func main() {

	//insertOne("a", 17)
	//insertMuti()
	//first()
	findAll()
}
