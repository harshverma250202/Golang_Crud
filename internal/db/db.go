package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB

	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "crud"
)

func init() {
	fmt.Println("Connecting to database")
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	Client, err = sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	_, err = Client.Query("SELECT * FROM users")
	if err != nil {
		_,err=Client.Query("CREATE TABLE users (id  SERIAL PRIMARY KEY,name varchar(255) NOT NULL,email varchar(255),password varchar(255));")
		if err!=nil{
			panic(err)
		}
	}

	log.Println("database successfully configured")
}
