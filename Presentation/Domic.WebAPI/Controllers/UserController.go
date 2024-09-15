package WebAPIController

import (
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

	createCommand := UseCaseUserCommand.CreateCommand{
		FirstName: "حسن",
		LastName:  "کرمی محب",
		Username:  "hasan_karami_moheb",
		Password:  "123456",
		Email:     "hasan_karami_moheb@gmail.com",
	}

	unitOfWork := InfrastructureConcrete.NewUnitOfWork(userController.db)

	createUserCommand := UseCaseUserCommand.NewCreateCommandHandler(
		InfrastructureConcrete.NewGlobalIdentityGenerator(),
		unitOfWork,
		InfrastructureConcrete.NewUserRepository(unitOfWork.GetTransaction()),
		InfrastructureConcrete.NewEventRepository(unitOfWork.GetTransaction()),
	)

	commandResult := createUserCommand.Handle(&createCommand)

	if !commandResult.Result {
		return context.String(http.StatusOK, "Error")
	}

	return context.String(http.StatusOK, "Successfully Created!")

}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}
