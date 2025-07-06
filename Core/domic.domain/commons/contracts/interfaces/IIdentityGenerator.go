package interfaces

type IIdentityGenerator interface {
	GetRandom(count byte) string
}
