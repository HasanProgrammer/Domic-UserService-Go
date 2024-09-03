package main

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Infrastructure/Concretes"
	"Domic.Persistence"
	"Domic.Persistence/Models"
	"Domic.UseCase/UserUseCase/Commands/Create"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()

	db := Persistence.NewSqlContext("sqlserver://sa:Domic123@127.0.0.1:1633?database=UserService").GetContext()

	db.AutoMigrate(&InfrastructureModel.EventModel{})
	db.AutoMigrate(&InfrastructureModel.UserModel{})

	go e.GET("user/create", func(c echo.Context) error {

		commandChannel := make(chan DomainCommonDTO.Results[bool])

		createCommand := UseCaseUserCommand.CreateCommand{
			FirstName: "حسن",
			LastName:  "کرمی محب",
			Username:  "hasan_karami_moheb",
			Password:  "123456",
			Email:     "hasan_karami_moheb@gmail.com",
		}

		unitOfWork := InfrastructureConcrete.NewUnitOfWork(db)

		createUserCommand := UseCaseUserCommand.NewCreateCommandHandler(
			unitOfWork,
			InfrastructureConcrete.NewUserRepository(unitOfWork.Transaction()),
			InfrastructureConcrete.NewEventRepository(unitOfWork.Transaction()),
		)

		go createUserCommand.Handle(&createCommand, commandChannel)

		commandResult := <-commandChannel

		if !commandResult.Result {
			return c.String(http.StatusOK, "Error")
		}

		return c.String(http.StatusOK, "Successfully Created!")

	})

	e.Logger.Fatal(e.Start(":1323"))

}
