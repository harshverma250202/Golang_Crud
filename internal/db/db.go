package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB

	// host     = "localhost"
	// port     = 5432
	// user     = "postgres"
	// password = "postgres"
	// dbname   = "crud"

)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	fmt.Println("Connecting to database")
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// var err error
	Client, err = sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	_, err = Client.Query("SELECT * FROM users")
	if err != nil {
		_, err = Client.Query("CREATE TABLE users (id  SERIAL PRIMARY KEY,name varchar(255) NOT NULL,email varchar(255),password varchar(255));")
		if err != nil {
			panic(err)
		}
	}

	log.Println("database successfully configured")
}
