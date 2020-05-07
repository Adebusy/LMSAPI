package utilities

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

//ConnectMe instantiate connection to database (*sql.DB, error)
func ConnectMe() {
	godotenv.Load()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		GoDotEnvVariable("Server"), GoDotEnvVariable("user"), GoDotEnvVariable("Password"), GoDotEnvVariable("Port"), GoDotEnvVariable("Database"))
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	} else {
		fmt.Println("no eerror")
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no eerror 2")
	}
	//return db, err
}
