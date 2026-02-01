package business_context_auth_service_model

type UserLoginResponseModel struct {
	AppUserId    string `json:"appUserId"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
