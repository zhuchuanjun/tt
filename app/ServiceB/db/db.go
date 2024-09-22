package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "test"
)

var psql *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	psql, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = psql.Ping()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

type User struct {
	Id    int64
	Name  string
	Email string
}

func GetById(id int64) *User {
	u := new(User)
	query := `select id, name, email from users where id = $1;`
	err := psql.QueryRow(query, id).Scan(&u.Id, &u.Name, &u.Email)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return u
}
