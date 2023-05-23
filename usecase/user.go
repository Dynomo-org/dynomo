package usecase

import (
	"context"
	"dynapgen/repository/db"
	"dynapgen/utils/log"
	"dynapgen/utils/tokenizer"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordSalt = 10
)

var (
	ErrorUserExists   = errors.New("user with the given email is already exists")
	ErrorUserNotFound = errors.New("user not found")
)

func (uc *Usecase) GetUserInfo(ctx context.Context, userID string) (UserInfo, error) {
	user, err := uc.db.GetUserByID(ctx, userID)
	if err != nil {
		log.Error(map[string]interface{}{"userID": userID}, err, "uc.get.GetUserByID() got error - GetUserInfo")
		return UserInfo{}, err
	}

	return UserInfo{
		ID:       user.ID,
		FullName: user.FullName,
	}, nil
}

func (uc *Usecase) LoginUser(ctx context.Context, user User) (AuthUserResponse, error) {
	savedUser, err := uc.db.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Error(user, err, "uc.db.GetUserByEmail() got error - LoginUser")
		return AuthUserResponse{}, err
	}
	if savedUser.ID == "" {
		return AuthUserResponse{}, ErrorUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return AuthUserResponse{}, ErrorUserNotFound
		}
		log.Error(user, err, "bcrypt.CompareHashAndPassword() got error - CompareHashAndPassword")
		return AuthUserResponse{}, err
	}

	token, err := tokenizer.GenerateJWTToken(map[string]interface{}{
		"id": savedUser.ID,
	})
	if err != nil {
		log.Error(savedUser, err, "tokenizer.GenerateJWTToken() got error - RegisterUser")
		return AuthUserResponse{}, err
	}

	return AuthUserResponse{
		Token: token,
		ID:    savedUser.ID,
	}, nil
}

func (uc *Usecase) RegisterUser(ctx context.Context, user User) (AuthUserResponse, error) {
	savedUser, err := uc.db.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Error(user, err, "uc.db.GetUserByEmail() got error - RegisterUser")
		return AuthUserResponse{}, err
	}
	if savedUser.ID != "" {
		return AuthUserResponse{}, ErrorUserExists
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), passwordSalt)
	if err != nil {
		log.Error(user, err, "bcrypt.GenerateFromPassword() got error - RegisterUser")
		return AuthUserResponse{}, err
	}

	userID := uuid.NewString()
	param := db.User{
		ID:       userID,
		Email:    user.Email,
		Password: string(hashedPassword),
		FullName: user.FullName,
	}

	err = uc.db.InsertUser(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.db.InsertUser() got error - RegisterUser")
		return AuthUserResponse{}, err
	}

	token, err := tokenizer.GenerateJWTToken(map[string]interface{}{
		"id": userID,
	})
	if err != nil {
		log.Error(param, err, "tokenizer.GenerateJWTToken() got error - RegisterUser")
		return AuthUserResponse{}, err
	}

	return AuthUserResponse{
		Token: token,
		ID:    userID,
	}, nil
}
