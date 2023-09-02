package main

// RequestResponse is the structure of response data for all HTTP requests
type RequestResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// SuccessResponse wraps a success message
func SuccessResponse(data interface{}) RequestResponse {
	ret := RequestResponse{}
	ret.Success = true
	ret.Data = data

	return ret
}

// ErrorResponse wraps an error message in the RequestResponse format
func ErrorResponse(message string) RequestResponse {
	ret := RequestResponse{}
	ret.Success = false
	ret.Error = message

	return ret
}
