package services

import (
	"errors"
	"log"
	"miniauth/models"
	"miniauth/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func (s *AuthService) Register(user *models.User) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)
	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.UserRepo.FindEmail(email)
	log.Println("LOGIN EMAIL:", email)
log.Println("DB EMAIL:", user.Email)
	log.Println("INPUT PASSWORD:", password)
log.Println("DB HASH:", user.Password)
	if err != nil {
		return nil, errors.New("invalid credntails")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil{
		return nil, errors.New("invalid password")
	}
	return user, nil
}