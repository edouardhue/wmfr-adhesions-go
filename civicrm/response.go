package civicrm

type StatusResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version"`
	Count        int `json:"count"`
	Id           int `json:"id"`
}

type Status interface {
	Success() bool
	GetErrorMessage() string
}

func (r *StatusResponse) Success() bool {
	return r.IsError == 0
}

func (r *StatusResponse) GetErrorMessage() string {
	return r.ErrorMessage
}

type ResponseError struct {
	Message string
}

func (e ResponseError) Error() string {
	return e.Message
}