package DomainCommonContract

type ICommandHandler[TResult any] interface {
	Handle() (TResult, error)
}
