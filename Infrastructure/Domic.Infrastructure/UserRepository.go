package DomicInfrastructure

import (
	"domic.domain/User/Entities"
	"time"
)

type UserRepository struct{}

func (userRepository *UserRepository) Add(User *Entities.User) bool {

	//something in here

	return true

}

func (userRepository *UserRepository) AddAsync(User *Entities.User, result chan bool) {

	//something in here

	channel := true

	time.Sleep(5 * time.Second)

	result <- channel

}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
