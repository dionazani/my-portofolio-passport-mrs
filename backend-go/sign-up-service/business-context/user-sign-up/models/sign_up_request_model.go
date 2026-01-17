package businessContext_model_signUp

// SignUp represents the sign_up table structure
type SignUpReqModel struct {
	ID          string  `json:"id"`
	Fullname    string  `json:"fullName"`
	SignUpFrom  string  `json:"signUpFrom"`  // e.g., 'WEB', 'APP'
	Email       *string `json:"email"`       // Pointer used because it can be NULL
	MobilePhone *string `json:"mobilePhone"` // Pointer used because it can be NULL
}
