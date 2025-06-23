package Create

type CreateCommand struct {
	FirstName string
	LastName  string
}

func CreateCommandHandler(command *CreateCommand) bool {

	return true
}
