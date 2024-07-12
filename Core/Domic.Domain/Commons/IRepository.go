package Commons

type IRepository[TEntity IEntity] interface {
	Add(Entity TEntity) bool
}
