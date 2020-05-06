package crudal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Adebusy/VisitorsManager/AppCode"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

//ConnectMe instantiate connection to database
func ConnectMe() (db *sql.DB, err error) {
	godotenv.Load()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		AppCode.GoDotEnvVariable("Server"), AppCode.GoDotEnvVariable("user"), AppCode.GoDotEnvVariable("Password"), AppCode.GoDotEnvVariable("Port"), AppCode.GoDotEnvVariable("Database"))
	db, err = sql.Open("sqlserver", connString)
	ctx := context.Background()
	err = db.PingContext(ctx)
	return db, err
}
