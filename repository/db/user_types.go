package db

type User struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	FullName string `db:"full_name"`
}

type UserInfo struct {
	ID       string `db:"id"`
	FullName string `db:"full_name"`
	RoleName string `db:"role_name"`
}
