package Commons

type ICommandHandler[TCommand ICommand[TResult], TResult any] interface {
	Handle(command TCommand[TResult]) TResult
}
