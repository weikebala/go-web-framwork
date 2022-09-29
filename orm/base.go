package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mygin/config"
)
var DB *gorm.DB

func init(){
	var err error
	DB, err = newDb()
	if err != nil {
		panic(err)
	}
}
func newDb() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.Db_dsn), &gorm.Config{})
}
