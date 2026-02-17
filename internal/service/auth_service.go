package service

import (
	"errors"

	"github.com/aryansehgal-tech/NotesApp/internal/models"
	"github.com/aryansehgal-tech/NotesApp/internal/repository"
	"github.com/aryansehgal-tech/NotesApp/internal/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(email, password string) error {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser.ID != 0 {
		return errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: hashedPassword,
	}

	return s.userRepo.Create(&user)
}
