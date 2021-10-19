package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint      `json:"id,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Name         string    `json:"Name,omitempty"`
	Username     string    `json:"Username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Picture      string    `json:"picture,omitempty"`
	Password     string    `json:"password,omitempty"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (u *User) HashPassword() error {
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(PasswordHash)
	return nil
}

func (u User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}
