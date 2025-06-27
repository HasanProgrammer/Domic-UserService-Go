package Interfaces

import (
	"domic.domain/Commons/Entities"
)

type IEventRepository interface {
	IRepository[string, *Entities.Event]
}
