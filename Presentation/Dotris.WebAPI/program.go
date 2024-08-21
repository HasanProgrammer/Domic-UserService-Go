package main

import (
	"Dotris.Infrastructure/Concretes/UserRepository"
	"Dotris.UseCase/UserUseCase/Commands/Create"
	"gorm.io/gorm"
)

func main() {

	createUserCommand :=
		UseCaseUserCommand.NewCreateCommandHandler(
			InfrastructureConcretesRepository.NewUserRepository(&gorm.DB{}),
		)

	createUserCommand.Handle()

}
