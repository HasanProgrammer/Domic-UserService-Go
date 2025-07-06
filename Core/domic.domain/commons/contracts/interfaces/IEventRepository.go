package contracts

import (
	"domic.domain/commons/Entities"
)

type IEventRepository interface {
	IRepository[string, *Entities.Event]
}
