package UseCaseUserCommand

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/Commons/DTOs"
	"time"
)

type SignInCommandHandler struct {
	jsonWebToken DomainCommonContract.IJsonWebToken
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

func NewSignInCommandHandler(jsonWebToken DomainCommonContract.IJsonWebToken) *SignInCommandHandler {
	return &SignInCommandHandler{jsonWebToken: jsonWebToken}
}
