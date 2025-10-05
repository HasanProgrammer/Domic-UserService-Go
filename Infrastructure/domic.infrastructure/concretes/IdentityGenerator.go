package concretes

type IdentityGenerator struct {
}

func (generator *IdentityGenerator) GetRandom(count byte) string {
	return ""
}

func NewIdentityGenerator() *IdentityGenerator {
	return &IdentityGenerator{}
}
