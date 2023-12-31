package injector

import (
	"ahsmha/Tasks/domain/repository"
	"ahsmha/Tasks/handler"
	"ahsmha/Tasks/infra"
	"ahsmha/Tasks/usecase"
)

func InjectTaskHandler() handler.TaskHandler {
	return handler.NewTaskHandler(InjectTaskUsecase())
}

func InjectTaskUsecase() usecase.TaskUsecase {
	TaskRepository := InjectTaskRepository()
	UserRepository := InjectUserRepository()

	return usecase.NewTaskUsecase(TaskRepository, UserRepository)
}

func InjectTaskRepository() repository.TaskRepository {
	sqlHandler := InjectDB()

	return infra.NewTaskRepository(sqlHandler)
}
