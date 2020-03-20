package storage

import "database/sql"

type mySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) Storage {
	return &mySQL{db: db}
}

func (m *mySQL) CreateUser(u *User) error {
	const q = `insert into users values (?,?,?)`
	_, err := m.db.Exec(q, u.ID, u.Name, u.Age)
	return err
}

func (m *mySQL) UpdateUser(u *User) error {
	return nil
}

func (m *mySQL) DeleteUser(id int) error {
	return nil
}

func (m *mySQL) GetUser(id int) (User, error) {
	return User{}, nil
}
