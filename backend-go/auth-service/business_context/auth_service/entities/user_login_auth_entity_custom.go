package business_context_auth_service_entity

type UserLoginAuthEntity struct {
	Email       string
	AppPassword string
	AppUserId   string
	IsLock      bool
}
