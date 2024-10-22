package convertor

import (
	"kwai/internal/user/domain"
	"kwai/internal/user/infra/persistence/po"
)

func UserPo2Entity(userPo *po.User) *domain.UserEntity {
	return &domain.UserEntity{
		Name: userPo.Name,
		Age:  int32(userPo.Age),
	}
}
