package Interfaces

type IRepository[TEntity IEntity] interface {
	Add(Entity TEntity) bool
	AddAsync(Entity TEntity, result chan bool)
}
