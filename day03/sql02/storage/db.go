package storage

import "database/sql"

type mySQL struct {
	db *sql.DB
}

func (m *mySQL) CreateUser(u *User) error {
	return nil
}

func (m *mySQL) UpdateUser(u *User) error {
	return nil
}

func (m *mySQL) DeleteUser(id int) error {
	return nil
}

func (m *mySQL) GetUser(id int) (User, error) {
	return nil
}
