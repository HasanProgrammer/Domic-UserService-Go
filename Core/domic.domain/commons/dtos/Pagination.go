package DTOs

type PaginationRequest struct {
	PageSize  int
	PageIndex int
}

type PaginationResponse[TResult interface{}] struct {
	PageSize  int
	PageIndex int
	TotalItem int64
	Items     []TResult
	HasNext   bool
	HasPrev   bool
}
