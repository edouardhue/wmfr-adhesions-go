package civicrm

type Response interface {
	Success() bool
	GetErrorMessage() string
}

type ResponseError struct {
	Message string
}

func (e ResponseError) Error() string {
	return e.Message
}