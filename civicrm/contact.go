package civicrm

type ContactQuery struct {
	EMail string `json:"email"`
}

type GetContactResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version"`
	Count        int `json:"count"`
	Id           int `json:"id"`
}

func (r *GetContactResponse) Success() bool {
	return r.IsError == 0
}

func (r *GetContactResponse) GetErrorMessage() string {
	return r.ErrorMessage
}