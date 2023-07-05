package injector

import (
	"ahsmha/notes/domain/repository"
	"ahsmha/notes/handler"
	"ahsmha/notes/infra"
	"ahsmha/notes/usecase"
)

// inject auth handler
func InjectAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(InjectUserUsecase())
}

func InjectUserUsecase() usecase.UserUsecase {
	userRepository := InjectUserRepository()

	return usecase.NewUserUsecase(userRepository)
}

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()

	return infra.NewUserRepository(sqlHandler)
}
