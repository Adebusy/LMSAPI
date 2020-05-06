package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Adebusy/VisitorsManager/AppCode"
	cntr "github.com/Adebusy/dataScienceAPI/controller"
	cr "github.com/Adebusy/dataScienceAPI/crudal"
	"github.com/Adebusy/dataScienceAPI/docs"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/celler/httputil"
)

var db *sql.DB
var err error

func init() {
	db, err = cr.ConnectMe()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no connection error")
	}
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
	r.POST("/user/CreateUser/", cntr.CreateUser)                                                  //done
	r.POST("/course/CreateCourse/", cntr.CreateCourse)                                            //done
	r.POST("/user/UpdateUserDetail/", cntr.UpdateUserDetail)                                      //done
	r.GET("/user/GetUserFullInfo/:EmailAddress", cntr.GetUserFullInfo)                            //done
	r.POST("/question/CreateNewQuestion/", cntr.CreateNewQuestion)                                //done
	r.GET("/question/FetchQuestionsByCourse/:StudentID/:CourseName", cntr.FetchQuestionsByCourse) //done
	r.Run(AppCode.GoDotEnvVariable("AppPort"))
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
