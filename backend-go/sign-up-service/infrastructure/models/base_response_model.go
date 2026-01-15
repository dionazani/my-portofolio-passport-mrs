package infrastructure_model

import "strconv"

// BaseResponse is the standard wrapper for all API responses
type BaseResponse struct {
	HTTPStatusCode string      `json:"httpStatusCode"`
	Status         string      `json:"status"`
	Timestamp      string      `json:"timestamp"`
	Data           interface{} `json:"data,omitempty"` // Can hold any type (User, etc.)
}

// In models/base_response.go
func (b *BaseResponse) GetIntStatusCode() int {
	code, _ := strconv.Atoi(b.HTTPStatusCode)
	return code
}
