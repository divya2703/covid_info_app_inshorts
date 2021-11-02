package errors

// swagger:model ServiceError
type ServiceError struct {

	//status code of the error
	StatusCode int `json:"status"`
	//business error message
	ErrorMessage string `json:"message"`
}
