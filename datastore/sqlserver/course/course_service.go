package course

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Adebusy/dataScienceAPI/driver/sqlserver"
	md "github.com/Adebusy/dataScienceAPI/modules"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type CourseService interface {
	Create(ctx context.Context, course md.Course) bool
	CheckIfCourseExistBool(ctx context.Context, course md.Course) bool
	CheckIfCourseExist(ctx context.Context, course md.Course) (md.TblCourse, error)
}

var db *sql.DB
var err error

var dbGorm *gorm.DB
var errGorm error
var respVal bool

type courseService struct {
	db     *sql.DB
	dbGorm *gorm.DB
}

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

//NewCourseService accessible from other packages
func NewCourseService(db *sql.DB, dbGorm *gorm.DB) CourseService {
	return &courseService{db, dbGorm}
}

//Create create new course
func (ts courseService) Create(ctx context.Context, course md.Course) bool {
	respVal = false
	objReq := md.TblCourse{CourseName: strings.ToUpper(course.CourseName), CourseCode: strings.ToUpper(course.CourseCode), CourseCategory: strings.ToUpper(course.CourseCategory), CourseStatus: "1", DateCreated: time.Now(), PassMark: course.PassMark}
	dbGorm.Table("tbl_course").Create(&objReq)
	if objReq.ID > 0 {
		respVal = true
	}
	return respVal
}

//CheckIfCourseExistBool validates course name and return true false
func (ts courseService) CheckIfCourseExistBool(ctx context.Context, courseNam md.Course) bool {
	var respVal bool = false
	fmt.Println(strings.ToUpper(courseNam.CourseName))
	objReq := md.TblCourse{}
	reqResp := dbGorm.Table("tbl_course").Where(`course_name =?`, strings.ToUpper(courseNam.CourseName)).First(&objReq).Error
	if reqResp == nil {
		fmt.Println("got here 1")
		if objReq.ID != 0 {
			fmt.Println(objReq.CourseCategory)
			respVal = true
		} else {
			respVal = false
		}
	} else {
		fmt.Println("got here 2")
		fmt.Println(reqResp)
		respVal = false
	}
	// fmt.Println(reqResp.Error.Error())
	//defer dbGorm.Close()
	return respVal
}

//CheckIfCourseExist validates course name and return true false
func (ts courseService) CheckIfCourseExist(ctx context.Context, course md.Course) (md.TblCourse, error) {
	var myCourse md.TblCourse
	reqResp := dbGorm.Table("tbl_course").Where(`course_name =?`, strings.ToUpper(course.CourseName)).First(&myCourse).Error
	if reqResp == nil {
		return myCourse, nil
	}
	//defer dbGorm.Close()
	return myCourse, reqResp
}
