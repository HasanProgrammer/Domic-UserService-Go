package DomainCommonContract

import DomainCommonDTO "Domic.Domain/Commons/DTOs"

type IJsonWebToken interface {
	Generate(claims map[string]interface{}) DomainCommonDTO.Result[string]
	Verify(tokenString string) DomainCommonDTO.Result[bool]
}
