package interfaces

import (
	"domic.domain/commons/entities"
)

type IEventRepository interface {
	IRepository[string, *entities.Event]
}
