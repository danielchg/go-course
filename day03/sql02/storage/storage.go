package storage

import (
	"errors"
	"os"
)

var (
	ErrDuplicate = errors.New("duplicate ID")
	ErrNotFound  = os.ErrNotExist
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Storage interface {
	CreateUser(*User) error
	UpdateUser(*User) error
	DeleteUser(int) error
	GetUser(int) (User, error)
}

var _ Storage = (*dummyStorage)(nil)

type dummyStorage struct {
	users []User
}

func (d *dummyStorage) CreateUser(u *User) error {
	for _, v := range d.users {
		if v.ID == u.ID {
			return ErrDuplicate
		}
	}
	d.users = append(d.users, *u)
	return nil
}

func (d *dummyStorage) UpdateUser(u *User) error {
	for i, v := range d.users {
		if v.ID == u.ID {
			d.users[i] = *u
			return nil
		}
	}
	return ErrNotFound
}

func (d *dummyStorage) DeleteUser(id int) error {
	for i, v := range d.users {
		if v.ID == id {
			if len(d.users) != 1 {
				d.users[i] = d.users[len(d.users)-1]
			}
			d.users = d.users[:len(d.users)-1]
			return nil
		}
	}
	return ErrNotFound
}

func (d *dummyStorage) GetUser(id int) (User, error) {
	for _, v := range d.users {
		if v.ID == id {
			return v, nil
		}
	}
	return User{}, ErrNotFound
}

// NewDummyStorage return a new dummystorage
func NewDummyStorage() Storage {
	return &dummyStorage{}
}
