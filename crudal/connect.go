package crudal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Adebusy/dataScienceAPI/utilities"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

//ConnectMe instantiate connection to database
func ConnectMe() (db *sql.DB, err error) {
	godotenv.Load()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		utilities.GoDotEnvVariable("Server"), utilities.GoDotEnvVariable("user"), utilities.GoDotEnvVariable("Password"), utilities.GoDotEnvVariable("Port"), utilities.GoDotEnvVariable("Database"))
	db, err = sql.Open("sqlserver", connString)
	ctx := context.Background()
	err = db.PingContext(ctx)
	return db, err
}
