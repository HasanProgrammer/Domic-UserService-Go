package main

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Infrastructure/Concretes"
	"Domic.Persistence/Models"
	"Domic.UseCase/UserUseCase/Commands/Create"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	e := echo.New()

	go e.GET("user/create", func(c echo.Context) error {

		commandChannel := make(chan DomainCommonDTO.Results[bool])

		db, err := gorm.Open(sqlserver.Open("sqlserver://sa:Domic123@127.0.0.1:1633?database=UserService"), &gorm.Config{})

		if err != nil {

		}

		db.AutoMigrate(&InfrastructureModel.EventModel{})
		db.AutoMigrate(&InfrastructureModel.UserModel{})

		createCommand := UseCaseUserCommand.CreateCommand{
			FirstName: "حسن",
			LastName:  "کرمی محب",
			Username:  "hasan_karami_moheb",
			Password:  "123456",
			Email:     "hasan_karami_moheb@gmail.com",
		}

		unitOfWork := InfrastructureConcrete.NewUnitOfWork(db)
		transaction := unitOfWork.Transaction()

		createUserCommand := UseCaseUserCommand.NewCreateCommandHandler(
			unitOfWork,
			InfrastructureConcrete.NewUserRepository(transaction),
			InfrastructureConcrete.NewEventRepository(transaction),
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
