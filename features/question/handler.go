package question

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/swaggo/swag/example/celler/httputil"

	"github.com/Adebusy/dataScienceAPI/driver/sqlserver"
	"github.com/Adebusy/dataScienceAPI/modules"

	csr "github.com/Adebusy/dataScienceAPI/datastore/sqlserver/course"
	"github.com/Adebusy/dataScienceAPI/datastore/sqlserver/question"
	"github.com/Adebusy/dataScienceAPI/datastore/sqlserver/student"
	"github.com/gin-gonic/gin"
)

var db *sql.DB
var err error

var dbGorm *gorm.DB
var errGorm error

type questionService struct {
	db     *sql.DB
	dbGorm *gorm.DB
}

var Newcourse = csr.NewCourseService(db, dbGorm)
var IQuestion = question.NewQuestionService(db, dbGorm)
var IStudent = student.NewstudentService(db)

func init() {
	dbGorm, errGorm = sqlserver.ConnectGorm()
	db, err = sqlserver.ConnectOLEDBC()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no connection error")
	}
	// dbGorm, errGorm := cr.ConnectGorm() //gorm.Open("mssql", "sqlserver://dbuser:Sterling123@sterlingazuredb.database.windows.net?database=newedupaydb")
	if errGorm != nil {
		fmt.Printf(errGorm.Error())
	} else {

	}
}

// CreateNewQuestion godoc
// @Summary create new question
// @Produce json
// @Param user body modules.Question true "create new question"
// @Success 200 {object} question.RequestResponse
// @Router /question/CreateNewQuestion/ [post]
func CreateNewQuestion(ctx *gin.Context) {
	var questionOBJ modules.Question
	var resp RequestResponse
	resp.ResponseCode = ""
	resp.ResponseMessage = ""
	var course modules.Course
	if err := ctx.ShouldBindJSON(&questionOBJ); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	course.CourseName = strings.ToUpper(questionOBJ.CourseName)
	fmt.Println("coursename is " + course.CourseName)
	//validate Request
	validateRequestResp := ValidateQuestionReq(questionOBJ)
	if validateRequestResp.ResponseCode != "" {
		ctx.JSON(http.StatusBadRequest, validateRequestResp)
		return
	}
	//check course name
	CourseNameValidation, _ := Newcourse.CheckIfCourseExist(ctx, course)
	if CourseNameValidation.CourseCode == "" {
		//if CourseNameValidation := Newcourse.CheckIfCourseExistBool(ctx, course); CourseNameValidation == false {
		resp.ResponseCode = "01"
		resp.ResponseMessage = "Course name supplied does not exist."
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	questionOBJ.CourseCode = CourseNameValidation.CourseCode

	//check if this question already exit for course
	if !IQuestion.CheckQuestionAlreadyExistForCourse(ctx, questionOBJ) {
		// create question for the course
		if IQuestion.CreateQuestion(ctx, questionOBJ) {
			resp.ResponseCode = "00"
			resp.ResponseMessage = fmt.Sprintf("Huray!!! question created successfully for %s", questionOBJ.CourseName)
			ctx.JSON(http.StatusOK, resp)
		} else {
			resp.ResponseCode = "01"
			resp.ResponseMessage = fmt.Sprintf("Unable to create question at the moment. Please try again later!!!")
			ctx.JSON(http.StatusOK, resp)
		}
	} else {
		resp.ResponseCode = "01"
		resp.ResponseMessage = fmt.Sprintf("This question has already been created for this course")
		ctx.JSON(http.StatusOK, resp)
	}
}

// FetchQuestionsByCourse godoc
// @Summary Fetch question for quis
// @Produce json
// @Param StudentID path string true "pass email as StudentID and course name as coursename"
// @Param CourseName path string true "courseName as course name"
// @Success 200 {object} question.QuisRequest
// @Router /question/FetchQuestionsByCourse/{StudentID}/{CourseName} [get]
func FetchQuestionsByCourse(ctx *gin.Context) {
	var resp RequestResponse
	resp.ResponseCode = ""
	resp.ResponseMessage = ""
	if courserName := ctx.Param("CourseName"); courserName == "" {
		resp.ResponseCode = "01"
		resp.ResponseMessage = "CourseName is required"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if courserName := ctx.Param("StudentID"); courserName == "" {
		resp.ResponseCode = "01"
		resp.ResponseMessage = "StudentID is required"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	//validate email address supplied
	if checkEmail := ValidateEmail(ctx.Param("StudentID")); checkEmail == false {
		resp.ResponseCode = "01"
		resp.ResponseMessage = "Student Email address must be valid."
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	//use email fetch student details
	doStudentCheck, err := IStudent.GetStudentByEmailAddress(ctx, strings.ToUpper(ctx.Param("StudentID")))

	if err != nil {
		fmt.Println(err.Error)
	}
	var cousesObj modules.Course
	cousesObj.CourseName = strings.ToUpper(ctx.Param("CourseName"))

	var requestObj modules.Question
	requestObj.CourseName = strings.ToUpper(ctx.Param("CourseName"))

	//use code code fetch course detail
	doCheck, _ := Newcourse.CheckIfCourseExist(ctx, cousesObj)
	//confirm course name exist
	if doCheck.CourseCode == "" {
		resp.ResponseCode = "02"
		resp.ResponseMessage = "Course does not exist."
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	//fetchQuestion use course name and maxinum exam question per quis for this course
	myquestion := IQuestion.GetQuestion(ctx, requestObj)
	var res QuisRequest
	res.MyQuestion = myquestion
	res.CourseDetails = doCheck
	res.StudentDetails = doStudentCheck
	ctx.JSON(http.StatusOK, res)
}

// TestResult godoc
// @Summary Keeps histories of test taken by registered student
// @Produce json
// @Param user body modules.TestResult true "Keeps histories of test taken by registered students"
// @Success 200 {object} modules.RequestResponse
// @Router /question/TestResult/ [post]
func TestResult(ctx *gin.Context) {
	var resqBody modules.TestResult
	var respBody modules.RequestResponse
	var course modules.Course
	course.CourseName = resqBody.TestID
	err := ctx.ShouldBindJSON(&resqBody)
	if err != nil {
		fmt.Println(err.Error)
	}
	//valdate student id
	_, erro := IStudent.GetStudentByEmailAddress(ctx, resqBody.StudentID)
	if erro != nil {
		fmt.Println(err.Error)
		respBody.ResponseCode = "01"
		respBody.ResponseMessage = "Unable to validate student email address at the moment. Please try again later!!"
		ctx.JSON(http.StatusOK, respBody)
		return
	}
	//confirm that course exist
	if checkCourse := Newcourse.CheckIfCourseExistBool(ctx, course); checkCourse == false {
		respBody.ResponseCode = "01"
		respBody.ResponseMessage = "Test name does not exist. Please re-confirm and try again!!"
		ctx.JSON(http.StatusOK, respBody)
		return
	}
	//do insert into tbl_test table
	doinsert := IQuestion.SubmitTestResult(ctx, resqBody)
	if doinsert <= 0 {
		respBody.ResponseCode = "02"
		respBody.ResponseMessage = "Unable to save result at the moment, please try again later."
	} else {
		respBody.ResponseCode = "00"
		respBody.ResponseMessage = "Result submitted successfully."
	}
	ctx.JSON(http.StatusOK, respBody)
}
