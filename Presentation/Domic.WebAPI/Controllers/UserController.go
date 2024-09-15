package WebAPIController

import (
	"Domic.Infrastructure/Concretes"
	"Domic.Persistence"
	UserCreate "Domic.UseCase/UserUseCase/Commands/Create"
	UserSignIn "Domic.UseCase/UserUseCase/Commands/SignIn"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	db *gorm.DB
}

func (userController *UserController) Create(context echo.Context) error {

	createCommand := UserCreate.CreateCommand{
		FirstName: "حسن",
		LastName:  "کرمی محب",
		Username:  "hasan_karami_moheb",
		Password:  "123456",
		Email:     "hasan_karami_moheb@gmail.com",
	}

	unitOfWork := InfrastructureConcrete.NewUnitOfWork(userController.db)

	createUserCommand := UserCreate.NewCreateCommandHandler(
		InfrastructureConcrete.NewGlobalIdentityGenerator(),
		unitOfWork,
		InfrastructureConcrete.NewUserRepository(unitOfWork.GetTransaction()),
		InfrastructureConcrete.NewEventRepository(unitOfWork.GetTransaction()),
	)

	commandResult := createUserCommand.Handle(&createCommand)

	if len(commandResult.Errors) > 0 {

		result, err := InfrastructureConcrete.NewSerializer().Serialize(commandResult.Errors)

		if err != nil {
		}

		return context.String(http.StatusOK, result)

	}

	return context.String(http.StatusOK, "Successfully Operation")

}

func (userController *UserController) SignIn(context echo.Context) error {

	signInCommand := UserSignIn.SignInCommand{
		Username: "",
		Password: "",
	}

	commandResult := UserSignIn.NewSignInCommandHandler(InfrastructureConcrete.NewJsonWebToken()).Handle(&signInCommand)

	if len(commandResult.Errors) > 0 {

		result, err := InfrastructureConcrete.NewSerializer().Serialize(commandResult.Errors)

		if err != nil {
		}

		return context.String(http.StatusOK, result)
	}

	return context.String(http.StatusOK, "Successfully Operation")

}

func NewUserController() *UserController {
	return &UserController{
		db: Persistence.NewSqlContext("sqlserver://sa:Domic123@127.0.0.1:1633?database=UserService").GetContext(),
	}
}
