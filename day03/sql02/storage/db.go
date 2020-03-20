package storage

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type mySQL struct {
	db *sql.DB
}

//NewMySQL create DB
func NewMySQL(db *sql.DB) Storage {
	return &mySQL{db: db}
}

func (m *mySQL) CreateUser(u *User) error {
	const q = `insert into users values (?,?,?)`
	_, err := m.db.Exec(q, u.ID, u.Name, u.Age)
	if mErr, ok := err.(*mysql.MySQLError); ok {
		if mErr.Number == 1062 {
			return ErrDuplicate
		}
	}
	return err
}

func (m *mySQL) UpdateUser(u *User) error {
	const q = `update users (name,age) values (?,?) where ID = ?`
	result, err := m.db.Exec(q, u.Name, u.Age, u.ID)
	if err != nil {
		return err
	}
	var n int64
	n, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

func (m *mySQL) DeleteUser(id int) error {
	return nil
}

func (m *mySQL) GetUser(id int) (User, error) {
	return User{}, nil
}
