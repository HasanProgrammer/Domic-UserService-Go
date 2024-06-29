package Entities

type User struct {
	//private
	firstName string
	lastName  string

	//public
	FirstName string
	LastName  string
}

// AddEvent inherit of <IEntity>
func (user User) AddEvent() {

}

func NewUser(FirstName string, LastName string) *User {
	return &User{FirstName: FirstName, LastName: LastName}
}
