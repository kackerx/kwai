package user

import (
	"fmt"

	db2 "kwai/internal/common/db"
	"kwai/internal/user/config"
	"kwai/internal/user/domain"
	"kwai/internal/user/infra/impl"
	"kwai/internal/user/infra/persistence/dal"
)

type Application struct {
	userDomain *domain.UserDomainService
}

func NewApp(basename string) *Application {
	cfg := config.New()
	fmt.Println(cfg)

	dsn := "root:Wasd4044@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	err := db2.Init(dsn)
	if err != nil {
		panic(err)
	}

	DB := db2.GetDB()
	data := dal.NewData(DB)
	userRepo := impl.NewUserRepo(data)
	userDomain := domain.NewUserDomainService(userRepo)

	return &Application{userDomain}
}

func (a *Application) GetUser(uid string) *domain.UserEntity {
	return a.userDomain.GetUser(uid)
}
