package DomainCommonContract

type IUnitOfWork interface {
	CommitTransaction() error
	RollbackTransaction() error
}
