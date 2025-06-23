package Helpers

type JsonBody struct{}

type JsonResponse struct {
	Code    int32
	Message string
	Body    JsonBody
}
