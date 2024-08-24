package DomainCommonDTO

type Result[TOutPut any] struct {
	e      error
	OutPut TOutPut
}
