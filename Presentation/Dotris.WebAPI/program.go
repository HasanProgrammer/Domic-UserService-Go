package main

import (
	InfrastructureConcrete "Domic.Infrastructure/Concretes"
	"gorm.io/gorm"
)

func main() {

	createCommand := UseCaseUserCommand.CreateCommand{}
	unitOfWork := InfrastructureConcrete.NewUnitOfWork(&gorm.DB{})
	userRepository := InfrastructureConcreteRepository.NewUserRepository(unitOfWork.Transaction())

	createUserCommand := UseCaseUserCommand.NewCreateCommandHandler(&createCommand, unitOfWork, userRepository)

	createUserCommand.Handle()

}
