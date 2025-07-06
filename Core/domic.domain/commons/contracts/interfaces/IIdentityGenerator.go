package contracts

type IIdentityGenerator interface {
	GetRandom(count byte) string
}
