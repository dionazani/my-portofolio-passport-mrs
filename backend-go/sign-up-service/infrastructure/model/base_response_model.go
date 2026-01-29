package infrastructure_model

import "strconv"

// BaseResponseModel is the standard wrapper for all API responses
type BaseResponseModel struct {
	HTTPStatusCode string      `json:"httpStatusCode"`
	Status         string      `json:"status"`
	Timestamp      string      `json:"timestamp"`
	Data           interface{} `json:"data,omitempty"` // Can hold any type (User, etc.)
}

// In model/base_response.go
func (b *BaseResponseModel) GetIntStatusCode() int {
	code, _ := strconv.Atoi(b.HTTPStatusCode)
	return code
}
