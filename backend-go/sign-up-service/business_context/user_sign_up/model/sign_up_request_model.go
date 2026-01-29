package business_context_user_sign_up_model

// SignUp represents the sign_up table structure
type SignUpRequestModel struct {
	ID          string  `json:"id"`
	Fullname    string  `json:"fullName"`
	SignUpFrom  string  `json:"signUpFrom"`  // e.g., 'WEB', 'APP'
	Email       *string `json:"email"`       // Pointer used because it can be NULL
	MobilePhone *string `json:"mobilePhone"` // Pointer used because it can be NULL
	Password    string  `json:"password"`
}
