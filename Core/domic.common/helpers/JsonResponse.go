package helpers

type JsonBody interface{}

type JsonResponse struct {
	Code    int32
	Message string
	Body    JsonBody
}
