package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@/test")
	age := 27
	q := `
select
	id, name
from
	users
where 
	age <= ?
;
	
`
	stmt, err := db.Prepare(`insert into users values (?,?,?)`)
	if err != nil {
		log.Fatal(err)
	}

	for i, r := range [][3]interface{}{
		{2, "Vito", 2},
		{3, "Ugo", 2},
	} {
		result, err := stmt.Exec(r[:]...)
		if err != nil {
			log.Println("Cannot insert", i, err)
			continue
		}
		log.Println(result.RowsAffected())
	}

	rows, err := db.Query(q, age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %s\n", id, name)
	}
}
