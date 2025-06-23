package DTOs

type Results[TResult any] struct {
	Errors []error
	Result TResult
}
