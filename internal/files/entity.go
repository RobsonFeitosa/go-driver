package files

import (
	"errors"
	"time"
)

var (
	ErrOwnerRequire = errors.New("owner is required")

	ErrNameRequire = errors.New("name is required")

	ErrTypeRequire = errors.New("type is required")

	ErrPathRequire = errors.New("path is required")
)

type File struct {
	ID         int64     `json:"id"`
	FolderID   int64     `json:"-"`
	OwnerID    int64     `json:"owner_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *File) Validate() error {

	if f.OwnerID == 0 {
		return ErrOwnerRequire
	}

	if f.Name == "" {
		return ErrNameRequire
	}

	if f.Type == "" {
		return ErrTypeRequire
	}

	if f.Path == "" {
		return ErrPathRequire
	}

	return nil
}
