package injector

import (
	"ahsmha/notes/domain/repository"
	"ahsmha/notes/infra"
	"ahsmha/notes/usecase"
)

// inject auth handler first
func InjectUserUsecase() usecase.UserUsecase {
	userRepository := InjectUserRepository()

	return usecase.NewUserUsecase(userRepository)
}

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()

	return infra.NewUserRepository(sqlHandler)
}
