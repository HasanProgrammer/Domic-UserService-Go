package Domic_Common

type JsonBody struct{}

type JsonResponse struct {
	Code    int32
	Message string
	Body    JsonBody
}
