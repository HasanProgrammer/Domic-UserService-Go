package main

import (
	"Domic.Infrastructure/Concretes"
	"Domic.Infrastructure/Concretes/UserRepository"
	"Domic.UseCase/UserUseCase/Commands/Create"
	"gorm.io/gorm"
)

func main() {

	createCommand := UseCaseUserCommand.CreateCommand{}
	unitOfWork := InfrastructureConcrete.NewUnitOfWork(&gorm.DB{})
	userRepository := InfrastructureConcreteRepository.NewUserRepository(unitOfWork.Transaction())

	createUserCommand := UseCaseUserCommand.NewCreateCommandHandler(&createCommand, unitOfWork, userRepository)

	createUserCommand.Handle()

}
