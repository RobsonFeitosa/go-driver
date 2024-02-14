package folders

import (
	"errors"
	"time"
)

var (
	ErrNameRequire = errors.New("Name is require")
)

func New(name string, parentId int64) (*Folder, error) {
	f := Folder{
		ParentId: parentId,
		Name:     name,
	}

	err := f.Validate()
	if err != nil {
		return nil, err
	}

	return &f, nil
}

type Folder struct {
	ID         int64     `json:"id"`
	ParentId   int64     `json:"parent_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *Folder) Validate() error {
	if f.Name == "" {
		return ErrNameRequire
	}
}
