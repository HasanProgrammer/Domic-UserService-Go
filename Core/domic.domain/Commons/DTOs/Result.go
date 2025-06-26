package DTOs

type Result[TResult any] struct {
	Errors []error
	Result TResult
}
