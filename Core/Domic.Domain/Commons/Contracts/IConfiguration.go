package DomainCommonContract

type IConfiguration interface {
	GetConnectionString(key string) (string, error)
}
