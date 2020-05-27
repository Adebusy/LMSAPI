package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

//ConnectGorm connection string for gorm
func ConnectGorm() (*gorm.DB, error) {
	godotenv.Load()
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", GoDotEnvVariable("user"), GoDotEnvVariable("Password"), GoDotEnvVariable("Server"), GoDotEnvVariable("Database"))
	db, errGorm := gorm.Open(GoDotEnvVariable("client"), connectionString)
	if errGorm != nil {
		//log.Panic(errGorm.Error())
		log.Println(errGorm.Error())
	}
	//return &Client{: errGorm}
	return db, errGorm
}

//Client connect to diff databases
type Client struct {
	//conn *pgx.Conn
	dbGorm *gorm.DB
	dbSQL  *gorm.Errors
}

var db *sql.DB
var dbg *gorm.DB
var err error

//ConnectOLEDBC instantiate connection to database (*sql.DB, error) func ConnectOLEDBC() *Client {
func ConnectOLEDBC() (*sql.DB, error) {
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
	return db, err
	//return &Client{dbSQL: db}
}

//Close closes all connections
//func (c *Client) Close() error {
//return c.dbSQL.Close()
//}

//CloseGorm closes all connections
func (c *Client) CloseGorm() error {
	return c.dbGorm.Close()
}

// func (c *Client) Begin() error {
// 	return c.db.Debug().First(&modules.Track{}).Error
// }

// func (c *Client) End() error {
// 	return c.db.Commit().Error
// }

// func (c *Client) SetMaxConn(limit int) {
// 	c.db.DB().SetMaxOpenConns(limit)
// }

// func (c *Client) Stats() interface{} {
// 	return c.db.DB().Stats()
// }

//GoDotEnvVariable load env file
func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
