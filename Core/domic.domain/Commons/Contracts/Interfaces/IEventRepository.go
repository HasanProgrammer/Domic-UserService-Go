package Interfaces

import (
	"domic.domain/Commons/Entities"
)

type IEventRepository[TIdentity any] interface {
	IRepository[TIdentity, *Entities.Event[TIdentity]]
}
