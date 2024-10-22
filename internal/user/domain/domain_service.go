package domain

type UserDomainService struct {
	userRepo UserRepository
}

func NewUserDomainService(userRepo UserRepository) *UserDomainService {
	return &UserDomainService{userRepo: userRepo}
}

func (d *UserDomainService) GetUser(uid string) *UserEntity {
	return d.userRepo.GetUserByID(uid)
}
