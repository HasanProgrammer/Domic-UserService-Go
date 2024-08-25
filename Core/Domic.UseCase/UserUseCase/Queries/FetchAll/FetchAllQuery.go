package UseCaseUserQuery

import (
	UseCaseCommonDTO "Domic.Domain/Commons/DTOs"
)

type FetchAllQuery struct {
	UseCaseCommonDTO.PaginationRequest
}
