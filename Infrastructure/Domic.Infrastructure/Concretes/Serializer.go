package InfrastructureConcrete

import (
	"encoding/json"
)

type Serializer struct{}

func (serializer *Serializer) Serialize(object interface{}) (string, error) {

	bytes, err := json.Marshal(object)

	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func (serializer *Serializer) Deserialize(stringifyPayload string, targetObject interface{}) error {

	bytes := []byte(stringifyPayload)

	return json.Unmarshal(bytes, targetObject)

}

func NewSerializer() *Serializer {
	return &Serializer{}
}
