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
	const q = `insert into users (id, name, age) values (?,?,?)`
	_, err := m.db.Exec(q, u.ID, u.Name, u.Age)
	if mErr, ok := err.(*mysql.MySQLError); ok {
		if mErr.Number == 1062 {
			return ErrDuplicate
		}
	}
	return err
}

func (m *mySQL) UpdateUser(u *User) error {
	const q = `update users set name=?, age=? where ID = ?`
	r, err := m.db.Exec(q, u.Name, u.Age, u.ID)
	if err != nil {
		return err
	}
	var n int64
	n, err = r.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

func (m *mySQL) DeleteUser(id int) error {
	const q = `delete from users where id = ?`
	r, err := m.db.Exec(q, id)
	if err != nil {
		return err
	}
	var n int64

	n, err = r.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}

	return nil
}

func (m *mySQL) GetUser(id int) (User, error) {
	const q = `select id, name, age from users where id = ?`
	r := m.db.QueryRow(q, id)

	u := User{}
	if err := r.Scan(&u.ID, &u.Name, &u.Age); err != nil {
		if err == sql.ErrNoRows {
			return User{}, ErrNotFound
		}

		return User{}, err
	}

	return u, nil
}
