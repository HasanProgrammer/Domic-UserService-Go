package Interfaces

type IIdentityGenerator interface {
	GetRandom(count byte) string
}
