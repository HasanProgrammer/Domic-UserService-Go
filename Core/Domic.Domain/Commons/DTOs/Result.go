package DomainCommonDTO

type Result[TResult any] struct {
	Error  error
	Result TResult
}

type Results[TResult any] struct {
	Errors []error
	Result TResult
}
