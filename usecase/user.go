package usecase

import (
	"context"
	"dynapgen/constants"
	"dynapgen/repository/db"
	"dynapgen/util/log"
	"dynapgen/util/tokenizer"
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
	user, err := uc.db.GetUserInfoByID(ctx, userID)
	if err != nil {
		log.Error(err, "uc.db.GetUserInfoByID() got error - GetUserInfo", map[string]interface{}{"user_id": userID})
		return UserInfo{}, err
	}
	if user.ID == "" {
		return UserInfo{}, ErrorUserNotFound
	}

	return UserInfo{
		ID:       user.ID,
		FullName: user.FullName,
		RoleName: user.RoleName,
	}, nil
}

func (uc *Usecase) GetUserRoleIDMapByUserID(ctx context.Context, userID string) (map[string]struct{}, error) {
	cachedUserRoles, err := uc.cache.GetUserRoleIDMapByUserID(ctx, userID)
	if err != nil {
		log.Error(err, "uc.cache.GetUserRoleIDMapByUserID() got error - GetUserRoleIDMapByUserID", map[string]interface{}{"user_id": userID})
	}
	if cachedUserRoles != nil {
		result := make(map[string]struct{})
		for k := range cachedUserRoles {
			result[string(k)] = struct{}{}
		}

		return result, nil
	}

	userRoles, err := uc.db.GetUserRoleIDsByUserID(ctx, userID)
	if err != nil {
		log.Error(err, "uc.db.GetUserRoleIDsByUserID() got error - GetUserRoleIDMapByUserID", map[string]interface{}{"user_id": userID})
		return nil, err
	}

	result := make(map[string]struct{})
	cacheParam := make(map[constants.UserRole]struct{})
	for _, role := range userRoles {
		result[role] = struct{}{}
		cacheParam[constants.UserRole(role)] = struct{}{}
	}

	err = uc.cache.SetUserRoleIDMapForUserID(ctx, userID, cacheParam)
	if err != nil {
		log.Error(err, "uc.cache.SetUserRoleIDMapForUserID() got error - GetUserRoleIDMapByUserID", map[string]interface{}{
			"user_id":     userID,
			"cache_param": cacheParam,
		})
	}

	return result, nil
}

func (uc *Usecase) LoginUser(ctx context.Context, user User) (AuthUserResponse, error) {
	savedUser, err := uc.db.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Error(err, "uc.db.GetUserByEmail() got error - LoginUser", user)
		return AuthUserResponse{}, err
	}
	if savedUser.ID == "" {
		return AuthUserResponse{}, ErrorUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return AuthUserResponse{}, ErrorUserNotFound
		}
		log.Error(err, "bcrypt.CompareHashAndPassword() got error - CompareHashAndPassword", user)
		return AuthUserResponse{}, err
	}

	token, err := tokenizer.GenerateJWTToken(map[string]interface{}{
		"id": savedUser.ID,
	})
	if err != nil {
		log.Error(err, "tokenizer.GenerateJWTToken() got error - RegisterUser", savedUser)
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
		log.Error(err, "uc.db.GetUserByEmail() got error - RegisterUser", user)
		return AuthUserResponse{}, err
	}
	if savedUser.ID != "" {
		return AuthUserResponse{}, ErrorUserExists
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), passwordSalt)
	if err != nil {
		log.Error(err, "bcrypt.GenerateFromPassword() got error - RegisterUser", user)
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
		log.Error(err, "uc.db.InsertUser() got error - RegisterUser", param)
		return AuthUserResponse{}, err
	}

	err = uc.db.InsertUserRole(ctx, userID, string(constants.RoleUser))
	if err != nil {
		log.Error(err, "uc.db.InsertUserRole() got error - RegisterUser", param)
		return AuthUserResponse{}, err
	}

	token, err := tokenizer.GenerateJWTToken(map[string]interface{}{
		"id": userID,
	})
	if err != nil {
		log.Error(err, "tokenizer.GenerateJWTToken() got error - RegisterUser", param)
		return AuthUserResponse{}, err
	}

	return AuthUserResponse{
		Token: token,
		ID:    userID,
	}, nil
}
