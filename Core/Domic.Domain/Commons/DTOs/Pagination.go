package DomainCommonDTO

type PaginationRequest struct {
	PageSize  int64
	PageIndex int64
}

type PaginationResponse[TResult interface{}] struct {
	PageSize  int64
	PageIndex int64
	TotalItem int64
	Items     []TResult
	HasNext   bool
	HasPrev   bool
}
