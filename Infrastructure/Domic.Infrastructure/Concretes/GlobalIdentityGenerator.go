package InfrastructureConcrete

import "github.com/google/uuid"

type GlobalIdentityGenerator struct{}

func (globalIdentityGenerator *GlobalIdentityGenerator) Generate() string {

	id, err := uuid.NewUUID()

	if err != nil {

	}

	return id.String()

}

func NewGlobalIdentityGenerator() *GlobalIdentityGenerator {
	return &GlobalIdentityGenerator{}
}
