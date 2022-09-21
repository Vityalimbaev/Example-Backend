package user

import (
	"context"
	"encoding/json"
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/Vityalimbaev/Example-Backend/internal/constants"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/repository/user"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type ServiceI interface {
	SignIn(name, password string) (string, string, error)
	SaveUser(user *entity.User) (*entity.User, error)
	GetUser(_ context.Context, userSearchParams *entity.UserSearchParams) ([]entity.User, error)
	RefreshUserSession(context context.Context, refreshToken string, userId int) (string, string, error)
}

type service struct {
	userRepo user.RepositoryI
}

func NewService(userRepo user.RepositoryI) *service {
	return &service{userRepo: userRepo}
}

func (s *service) SignIn(email, password string) (string, string, error) {

	users, err := s.userRepo.GetUsers(&entity.UserSearchParams{Email: email})

	if err != nil {
		return "", "", err
	}

	if len(users) == 0 {
		return "", "", exception.UnauthorizedError
	}

	if err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(password)); err != nil {
		return "", "", exception.UnauthorizedError
	}

	accessToken, err := s.GenerateAccessToken(users[0])
	refreshToken := s.GenerateRefreshToken(50)

	if err != nil {
		return "", "", exception.InternalError
	}

	err = s.userRepo.InsertUserSession(&entity.UserSession{
		RefreshToken: refreshToken,
		DateExpire:   time.Now().Add(72 * time.Hour).Unix(),
		UserId:       users[0].Id,
	})

	if err != nil {
		return "", "", exception.InternalError
	}

	return accessToken, refreshToken, nil
}

func (s *service) SaveUser(user *entity.User) (*entity.User, error) {
	if !user.IsValidForSave() {
		return nil, exception.BadRequest
	}

	passwordHash, err := s.GetPasswordHash(user.Password)

	if err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	user.Password = passwordHash

	id, err := s.userRepo.InsertUser(user)

	user.Id = id
	user.Password = ""

	return user, err
}

func (s *service) GetUser(_ context.Context, userSearchParams *entity.UserSearchParams) ([]entity.User, error) {
	users, err := s.userRepo.GetUsers(userSearchParams)

	for index := range users {
		users[index].Password = ""
	}

	return users, err
}

func (s *service) RefreshUserSession(_ context.Context, refreshToken string, userId int) (string, string, error) {
	session, err := s.userRepo.GetUserSession(userId)

	if err != nil {
		return "", "", nil
	}

	if time.Now().Unix() > session.DateExpire && refreshToken == session.RefreshToken {
		return "", "", exception.UnauthorizedError
	}

	accounts, err := s.userRepo.GetUsers(&entity.UserSearchParams{Id: userId})

	if err != nil || len(accounts) == 0 {
		return "", "", exception.UnauthorizedError
	}

	accessToken, err := s.GenerateAccessToken(accounts[0])
	newRefreshToken := s.GenerateRefreshToken(50)

	if err != nil {
		return "", "", nil
	}

	err = s.userRepo.InsertUserSession(&entity.UserSession{UserId: userId, RefreshToken: newRefreshToken, DateExpire: time.Now().Unix()})

	if err != nil {
		return "", "", nil
	}

	return accessToken, newRefreshToken, nil
}

func (s *service) GenerateAccessToken(user entity.User) (string, error) {
	conf := config.GetServerConfig()
	expirationTime := time.Now().Add(time.Duration(conf.TokenExpireDuration) * time.Hour)

	sub := map[string]interface{}{
		constants.CtxUserIdKey:       user.Id,
		constants.CtxActiveStatusKey: user.ActiveStatus,
		constants.CtxRoleIdKey:       user.RoleId,
	}

	subJson, _ := json.Marshal(sub)

	claims := jwt.StandardClaims{
		ExpiresAt: jwt.At(expirationTime),
		Subject:   string(subJson),
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	keyBytes := []byte(conf.TokenKey)

	tokenString, err := tokenObj.SignedString(keyBytes)

	if err != nil {
		logrus.Error(err)
		return "", exception.InternalError
	}

	return tokenString, nil
}

func (s *service) GenerateRefreshToken(n int) string {

	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func (s *service) GetPasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logrus.Error(err)
		return string(passwordHash), err
	}
	return string(passwordHash), err
}
