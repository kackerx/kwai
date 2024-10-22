package domain

type UserRepository interface {
	GetUserByID(uid string) *UserEntity
}
