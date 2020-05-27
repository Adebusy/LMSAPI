package course

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/Adebusy/dataScienceAPI/driver/sqlserver"
	"github.com/Adebusy/dataScienceAPI/modules"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"

	"github.com/Adebusy/dataScienceAPI/datastore/sqlserver/course"
)

var db *sql.DB
var dbGorm *gorm.DB
var err error
var doNewCouurseService = course.NewCourseService(db, dbGorm)
var errGorm error
var respVal bool

func init() {
	dbGorm, errGorm = sqlserver.ConnectGorm()
	db, err = sqlserver.ConnectOLEDBC()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no connection error")
	}
	if errGorm != nil {
		fmt.Printf(errGorm.Error())
	} else {

	}
}

// CreateCourse godoc
// @Summary create new course
// @Produce json
// @Param user body modules.Course true "Create new course"
// @Success 200 {object} modules.Message
// @Router /course/CreateCourse [post]
func CreateCourse(ctx *gin.Context) {
	var course modules.Course
	var msg modules.Message
	fmt.Println("read record4")
	ctx.ShouldBindJSON(&course)
	if len(course.CourseName) < 3 {
		msg.Message = "Course name length must be greated than 3."
		ctx.JSON(http.StatusBadRequest, msg)
	}
	fmt.Println("read record5")
	//check if course name already exist
	checkAlreadyExist, _ := doNewCouurseService.CheckIfCourseExist(ctx, course)
	if checkAlreadyExist.CourseCode == "" {
		n := rand.Int63()
		var getcourseCode = strings.ToUpper(course.CourseName[0:3]) + strconv.FormatInt(n, 2)
		course.CourseCode = getcourseCode
		if respCreate := doNewCouurseService.Create(ctx, course); respCreate == true {
			msg.Message = "New course with code '" + getcourseCode + "' created successfully."
			ctx.JSON(http.StatusOK, msg)
		} else {
			msg.Message = "Unable to create new course successfully"
			ctx.JSON(http.StatusOK, msg)
		}
	} else {
		msg.Message = "Name already exists." //s+ reqErr.Error()
		ctx.JSON(http.StatusOK, msg)
	}
}
