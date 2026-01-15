package businessContext_userSignUp_models

// UserRequest represents the incoming JSON payload for user operations
type SignUpReq struct {
	// Use 'json:"key"' to match the incoming API request keys
	ID          string `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilePhone"`
}
