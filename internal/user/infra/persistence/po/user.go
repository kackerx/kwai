package po

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

func (u User) TableName() string {
	return "t_blog_user"
}
