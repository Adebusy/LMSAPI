package crudal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Adebusy/dataScienceAPI/utilities"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

var dbGorm *sql.DB
var errGorm error

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

//ConnectGorm connection string for gorm
func ConnectGorm() (db *gorm.DB, err error) {
	dbGorm, errGorm := gorm.Open("mssql", "sqlserver://dbuser:Sterling123@sterlingazuredb.database.windows.net?database=newedupaydb")
	return dbGorm, errGorm
}
