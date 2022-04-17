package auth

// Authentication Struct

type AuthUser struct {
	UserId string
	Role   string
}

func (a *AuthUser) GetUserId() string {
	return a.UserId
}

func (a *AuthUser) GetRole() string {
	return a.Role
}
