package db

import (
	"context"
	"database/sql"
)

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var result User
	if err := r.db.GetContext(ctx, &result, queryGetUserByEmail, email); err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	return result, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (User, error) {
	var result User
	if err := r.db.GetContext(ctx, &result, queryGetUserByID, id); err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	return result, nil
}

func (r *Repository) InsertUser(ctx context.Context, user User) error {
	if _, err := r.db.QueryContext(ctx, queryInsertUser, user.ID, user.Email, user.Password, user.FullName); err != nil {
		return err
	}

	return nil
}
