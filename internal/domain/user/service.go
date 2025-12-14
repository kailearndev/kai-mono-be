package user

import (
	"errors"
	"kai-mono-be/pkg/jwt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(req RegisterUserDTO) (User, error)
	LoginUser(req LoginUserDTO) (string, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) RegisterUser(req RegisterUserDTO) (User, error) {
	existingUser, _ := s.repo.FindByEmail(req.Email)

	if existingUser.ID != uuid.Nil {
		return User{}, errors.New("email already exists")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hash),
		FullName:  req.FullName,
		Role:      req.Role,
		AvatarURL: req.AvatarURL,
	}
	if err := s.repo.CreateUser(&user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) LoginUser(req LoginUserDTO) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}
	// Here you would normally generate a JWT or session token
	token, err := jwt.GenerateToken(user.ID.String(), user.Email, time.Hour*24)
	if err != nil {
		return "", err
	}
	return token, nil
}
