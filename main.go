package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	cntr "github.com/Adebusy/dataScienceAPI/controller"
	cr "github.com/Adebusy/dataScienceAPI/crudal"
	"github.com/Adebusy/dataScienceAPI/docs"
	"github.com/Adebusy/dataScienceAPI/model"
	ut "github.com/Adebusy/dataScienceAPI/utilities"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/celler/httputil"
)

//git remote add origin https://github.com/Adebusy/dataScienceAPI.git
var db *sql.DB
var err error

var dbGorm *gorm.DB
var errGorm error

func init() {
	db, err = cr.ConnectMe()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no connection error")
	}
	dbGorm, errGorm := cr.ConnectGorm() //gorm.Open("mssql", "sqlserver://dbuser:Sterling123@sterlingazuredb.database.windows.net?database=newedupaydb")
	if errGorm != nil {
		fmt.Printf(errGorm.Error())
	} else {
		//fmt.Println("mo connect ")
		//dbGorm.Debug().DropTableIfExists(&model.TblTestNewResult{})
		//dbGorm.SingularTable(true)
		//dbGorm.Debug().AutoMigrate(&model.TblTestNewResult{})
	}
	insertTest := model.TblTestNewResult{DateTaken: time.Now().Local(), StudentID: "alaoh.adebusy@gmail.com", TestID: "1203", TestResult: "6004"}
	dbGorm.Create(&insertTest)
	//dbGorm.Last(&insertTest)
	//fmt.Println(insertTest.ID)

}

// @title Data Science Central API
// @version 1.0
// @description Data science competency check API.
// @termsOfService http://swagger.io/terms/
// @contact.name Alao ramon Adebisi
// @contact.email alao.adebusy@gmail.com
// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE
// @BasePath /api/v1
func main() {
	//router := mux.NewRouter()
	docs.SwaggerInfo.Title = "Data analysis API "
	docs.SwaggerInfo.Description = "This is a Data analysis management API.."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.Host = "localhost:8092"
	docs.SwaggerInfo.BasePath = "/" //v1
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/user/CreateUser/", cntr.CreateUser)
	r.POST("/course/CreateCourse/", cntr.CreateCourse)
	r.POST("/user/UpdateUserDetail/", cntr.UpdateUserDetail)
	r.GET("/user/GetUserFullInfo/:EmailAddress", cntr.GetUserFullInfo)
	r.POST("/question/CreateNewQuestion/", cntr.CreateNewQuestion)
	r.POST("/question/TestResult/", cntr.TestResult)
	r.GET("/question/FetchQuestionsByCourse/:StudentID/:CourseName", cntr.FetchQuestionsByCourse)
	r.Run(ut.GoDotEnvVariable("AppPort"))
}
func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
