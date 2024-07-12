package DomicInfrastructure

import (
	"domic.domain/User/Entities"
)

type UserRepository struct{}

func (userRepository *UserRepository) Add(User *Entities.User) bool {

	//something in here

	return true

}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
