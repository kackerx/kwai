package dal

import (
	"gorm.io/gorm"

	"kwai/internal/user/infra/persistence/po"
)

type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) *Data {
	return &Data{db: db}
}

func (d *Data) SelectUserByID(id string) *po.User {
	return &po.User{
		Name: "kacker",
		Age:  29,
	}
}
