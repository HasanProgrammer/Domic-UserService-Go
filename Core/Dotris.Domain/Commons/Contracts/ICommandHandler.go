package DomainCommonContract

type ICommandHandler interface {
	Handle() error
}
