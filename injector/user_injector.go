package injector

import (
	"ahsmha/Tasks/domain/repository"
	"ahsmha/Tasks/infra"
	"ahsmha/Tasks/usecase"
)

// inject auth handler
// func InjectAuthHandler() handler.AuthHandler {
// 	return handler.NewAuthHandler(InjectUserUsecase())
// }

func InjectUserUsecase() usecase.UserUsecase {
	userRepository := InjectUserRepository()

	return usecase.NewUserUsecase(userRepository)
}

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()

	return infra.NewUserRepository(sqlHandler)
}
