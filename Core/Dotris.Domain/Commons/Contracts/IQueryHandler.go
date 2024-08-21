package DomainCommonContract

type IQueryHandler[TResult interface{}] interface {
	Handle() (TResult, error)
}
