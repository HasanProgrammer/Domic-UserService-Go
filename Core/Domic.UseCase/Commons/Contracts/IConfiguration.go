package UseCaseCommonContract

type IConfiguration interface {
	GetConnectionString(key string) (string, error)
}
