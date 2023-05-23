package db

const (
	queryGetUserByEmail = `select * from users where email = $1`
	queryGetUserByID    = `select * from users where id = $1`
	queryInsertUser     = `insert into users (id, email, password, full_name) values ($1, $2, $3, $4)`
)
