package viewmodel

type ErrorMessage struct {
	Code   int32    `json:"code"`
	Errors []string `json:"errors"`
}
