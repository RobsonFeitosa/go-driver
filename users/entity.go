package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNameRequire = errors.New("Name is required")

	ErrLoginRequire = errors.New("Login is required")

	ErrPasswordRequire = errors.New("Passwords is required and can't be blank")

	ErrPasswordLength = errors.New("Password must have at least 6 characters")
)

func New(name, login, password string) (*User, error) {
	u := User{Name: name, Login: login, ModifiedAt: time.Now()}
	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequire
	}

	if len(password) < 6 {
		return ErrPasswordLength
	}

	u.Password = fmt.Sprintf("%x", (md5.Sum([]byte(password))))

	return nil
}

func (u *User) Validate() error {

	if u.Name == "" {
		return ErrNameRequire
	}

	if u.Login == "" {
		return ErrLoginRequire
	}

	BlankPassword := fmt.Sprintf("%x", md5.Sum([]byte("")))
	if u.Password == BlankPassword {
		return ErrPasswordRequire
	}

	return nil
}