package users

import (
	"time"
)

func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (h *handler) authenticate(login, password string) (*User, error) {
	stmt := `SELECT * FROM "users" WHERE login=$1 and password=$2`
	row := h.db.QueryRow(stmt, login, encPass(password))

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (h *handler) updateLastLogin(u *User) error {
	u.LastLogin = time.Now()

	return Update(h.db, u.ID, u)
}

func Authenticate(login, password string) (*User, error) {
	u, err := gh.authenticate(login, password)
	if err != nil {
		return nil, err
	}

	err = gh.updateLastLogin(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
