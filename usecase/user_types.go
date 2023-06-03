package usecase

type AuthUserResponse struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}

type User struct {
	ID       string
	Email    string
	Password string
	FullName string
}

type UserInfo struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	RoleName string `json:"role_name"`
}
