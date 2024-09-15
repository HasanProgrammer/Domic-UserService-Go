package UseCaseUserCommand

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.UseCase/Commons/Contracts"
	"time"
)

type SignInCommandHandler struct {
	jsonWebToken UseCaseCommonContract.IJsonWebToken
}

func (signInCommandHandler *SignInCommandHandler) Handle(command *SignInCommand) DomainCommonDTO.Results[string] {

	var errors []error

	token := signInCommandHandler.jsonWebToken.Generate(
		map[string]interface{}{
			"username": command.Username,
			"password": command.Password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	return DomainCommonDTO.Results[string]{
		Errors: append(errors, token.Error),
		Result: token.Result,
	}

}

func NewSignInCommandHandler(jsonWebToken UseCaseCommonContract.IJsonWebToken) *SignInCommandHandler {
	return &SignInCommandHandler{jsonWebToken: jsonWebToken}
}
