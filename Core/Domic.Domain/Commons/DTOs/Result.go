package DomainCommonDTO

type Result[TOutPut any] struct {
	Error  error
	OutPut TOutPut
}

type Results[TOutPut any] struct {
	Errors []error
	OutPut TOutPut
}
