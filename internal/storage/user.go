package storage

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var users = map[string]User{} // key = email

type User struct {
	Email    string
	Password string
}

func CreateUser(email, password string) error {
	if _, exists := users[email]; exists {
		return errors.New("User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users[email] = User{Email: email, Password: string(hashedPassword)}
	return nil
}

func GetUser(email string) (User, error) {
	user, exists := users[email]
	if !exists {
		return User{}, errors.New("User not found")
	}

	return user, nil
}
