package impl

import (
	"kwai/internal/user/domain"
	"kwai/internal/user/infra/persistence/convertor"
	"kwai/internal/user/infra/persistence/dal"
)

type UserRepoImpl struct {
	data *dal.Data
}

func NewUserRepo(data *dal.Data) *UserRepoImpl {
	return &UserRepoImpl{data: data}
}

func (u *UserRepoImpl) GetUserByID(uid string) *domain.UserEntity {
	userPo := u.data.SelectUserByID(uid)
	return convertor.UserPo2Entity(userPo)
}
