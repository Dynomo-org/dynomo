package db

const (
	queryGetUserByEmail  = `select * from users where email = $1`
	queryGetUserInfoByID = `
		select 
			"user".id,
			"user".full_name,
			role.name as role_name
		from users "user"
		inner join user_roles user_role on user_role.user_id = "user".id
		inner join roles role on role.id = user_role.role_id
		where "user".id = $1
	`
	queryGetUserRoleIDsByUserID = `select role_id from user_roles where user_id = $1`
	queryInsertUser             = `insert into users (id, email, password, full_name) values ($1, $2, $3, $4)`
	queryInsertUserRole         = `insert into user_roles (user_id, role_id) values ($1, $2)`
)
