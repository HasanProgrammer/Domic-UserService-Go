package WebAPIController

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Infrastructure/Concretes"
	"Domic.UseCase/UserUseCase/Commands/Create"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	db *gorm.DB
}

func (userController *UserController) Create(context echo.Context) error {

	commandChannel := make(chan DomainCommonDTO.Results[bool])

	createCommand := UseCaseUserCommand.CreateCommand{
		FirstName: "حسن",
		LastName:  "کرمی محب",
		Username:  "hasan_karami_moheb",
		Password:  "123456",
		Email:     "hasan_karami_moheb@gmail.com",
	}

	unitOfWork := InfrastructureConcrete.NewUnitOfWork(userController.db)

	createUserCommand := UseCaseUserCommand.NewCreateCommandHandler(
		unitOfWork,
		InfrastructureConcrete.NewUserRepository(unitOfWork.GetTransaction()),
		InfrastructureConcrete.NewEventRepository(unitOfWork.GetTransaction()),
	)

	go createUserCommand.Handle(&createCommand, commandChannel)

	commandResult := <-commandChannel

	if !commandResult.Result {
		return context.String(http.StatusOK, "Error")
	}

	return context.String(http.StatusOK, "Successfully Created!")

}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}
