package db

import (
	"context"
	"database/sql"
)

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result User
	if err := r.db.GetContext(ctx, &result, queryGetUserByEmail, email); err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	return result, nil
}

func (r *Repository) GetUserInfoByID(ctx context.Context, id string) (UserInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result UserInfo
	if err := r.db.GetContext(ctx, &result, queryGetUserInfoByID, id); err != nil && err != sql.ErrNoRows {
		return UserInfo{}, err
	}

	return result, nil
}

func (r *Repository) GetUserRoleIDsByUserID(ctx context.Context, userID string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result []string
	if err := r.db.SelectContext(ctx, &result, queryGetUserRoleIDsByUserID, userID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) InsertUser(ctx context.Context, user User) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	if _, err := r.db.QueryContext(ctx, queryInsertUser, user.ID, user.Email, user.Password, user.FullName); err != nil {
		return err
	}

	return nil
}

func (r *Repository) InsertUserRole(ctx context.Context, userID, roleID string) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	if _, err := r.db.QueryContext(ctx, queryInsertUserRole, userID, roleID); err != nil {
		return err
	}

	return nil
}
