package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/victorlazzaroli/microservicesTest/auth/config"
	"github.com/victorlazzaroli/microservicesTest/auth/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type usersServices struct {
	db  *gorm.DB
	env *config.Env
}

type UsersServicesI interface {
	Signup(name, email, password string) error
	Signin(email, password string) (string, error)
}

// Signin implements UsersServicesI.
func (s *usersServices) Signin(email string, password string) (string, error) {
	var user models.User
	{
		result := s.db.First(&user, "email = ?", email)

		if result.Error != nil {
			return "", fmt.Errorf("Invalid email or password")
		}
	}
	{
		result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if result != nil {
			return "", fmt.Errorf("Invalid email or password")
		}
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userData": user,
		"sub":      user.ID,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(s.env.JWTSecret))
}

func (s *usersServices) Signup(name, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return fmt.Errorf("Internal Server Error")
	}

	result := s.db.Create(&models.User{Name: name, Email: email, Password: string(hash)})

	if result.Error != nil {
		return fmt.Errorf("Internal Server Error")
	}

	return nil

}

func NewUsersService(db *gorm.DB, env *config.Env) UsersServicesI {
	service := &usersServices{
		db:  db,
		env: env,
	}

	return service
}
