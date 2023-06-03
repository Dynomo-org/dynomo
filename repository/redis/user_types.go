package redis

type UserInfo struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	RoleName string `json:"role_name"`
}
