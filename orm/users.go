package orm

import (
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Name string
	Age uint8
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = strings.ToUpper(u.Name)
	return
}

