package DomainCommonDTO

type Result[TOutPut any] struct {
	Error  error
	OutPut TOutPut
}
