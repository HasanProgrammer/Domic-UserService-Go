package WebAPIController

import (
	"Domic.Infrastructure/Concretes"
	"Domic.Persistence"
	UserCreate "Domic.UseCase/UserUseCase/Commands/Create"
	UserSignIn "Domic.UseCase/UserUseCase/Commands/SignIn"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	connectionString string
}

// Create godoc
// @Summary
// @Description
// @Tags         Users
// @Accept       json
// @Produce      json
// @param        FirstName body  string  true "firstName"
// @param        LastName  body  string  true "lastName"
// @param        Username  body  string  true "username"
// @param        Password  body  string  true "password"
// @param        Email     body  string  true "email"
// @Success      200  {object}  error
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v1/users [post]
func (controller *UserController) Create(c echo.Context) error {

	createCommand := UserCreate.CreateCommand{
		FirstName: c.FormValue("FirstName"),
		LastName:  c.FormValue("LastName"),
		Username:  c.FormValue("Username"),
		Password:  c.FormValue("Password"),
		Email:     c.FormValue("Email"),
	}

	db := Persistence.NewSqlContext("").GetContext()

	unitOfWork := InfrastructureConcrete.NewUnitOfWork(db)

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

		return c.String(http.StatusOK, result)

	}

	return c.String(http.StatusOK, "Successfully Operation")

}

func (controller *UserController) Update(c echo.Context) error {

	db := Persistence.NewSqlContext(controller.connectionString).GetContext()

	createCommand := UserCreate.CreateCommand{
		FirstName: "حسن",
		LastName:  "کرمی محب",
		Username:  "hasan_karami_moheb",
		Password:  "123456",
		Email:     "hasan_karami_moheb@gmail.com",
	}

	unitOfWork := InfrastructureConcrete.NewUnitOfWork(db)

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

		return c.String(http.StatusOK, result)

	}

	return c.String(http.StatusOK, "Successfully Operation")

}

func (controller *UserController) SignIn(c echo.Context) error {

	//db := Persistence.NewSqlContext(controller.connectionString).GetContext()

	signInCommand := UserSignIn.SignInCommand{
		Username: c.FormValue("Username"),
		Password: c.FormValue("Password"),
	}

	commandResult := UserSignIn.NewSignInCommandHandler(InfrastructureConcrete.NewJsonWebToken()).Handle(&signInCommand)

	if len(commandResult.Errors) > 0 {

		result, err := InfrastructureConcrete.NewSerializer().Serialize(commandResult.Errors)

		if err != nil {
		}

		return c.String(http.StatusOK, result)
	}

	return c.String(http.StatusOK, "Successfully Operation")

}

func NewUserController() *UserController {

	connectionString, err := InfrastructureConcrete.NewConfiguration().GetConnectionString("SqlServer")

	if err != nil {

	}

	return &UserController{
		connectionString: connectionString,
	}

}
