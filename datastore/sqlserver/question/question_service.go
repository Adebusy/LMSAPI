package question

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Adebusy/dataScienceAPI/driver/sqlserver"
	"github.com/Adebusy/dataScienceAPI/modules"
	md "github.com/Adebusy/dataScienceAPI/modules"
	"github.com/jinzhu/gorm"
)

var db *sql.DB
var err error

var dbGorm *gorm.DB
var errGorm error

type questionService struct {
	db     *sql.DB
	dbGorm *gorm.DB
}

var recsliceRaw = []*md.Questions{}
var randonNumbers int = 0

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

type QuestionService interface {
	CreateQuestion(ctx context.Context, questionObj md.Question) bool
	GetQuestion(ctx context.Context, requestObj md.Question) []*md.Questions
	SubmitTestResult(ctx context.Context, ts md.TestResult) int
	CheckQuestionAlreadyExistForCourse(ctx context.Context, questionObj md.Question) bool
}

//NewQuestionService accessible from other packages
func NewQuestionService(db *sql.DB, dbGorm *gorm.DB) QuestionService {
	return &questionService{db, dbGorm}
}

//CheckQuestionAlreadyExistForCourse confirm question
func (qts questionService) CheckQuestionAlreadyExistForCourse(ctx context.Context, questionObj md.Question) bool {
	createOBJ := modules.TblQuestion{}
	query := dbGorm.Table(`tbl_question`).Where(`course_name =? and course_code=? and question=?`, strings.ToUpper(questionObj.CourseName), strings.ToUpper(questionObj.CourseCode), strings.ToUpper(questionObj.Question)).First(&createOBJ).Error
	if query != nil {
		fmt.Println(query)
	}
	if createOBJ.ID != 0 {
		return true
	}
	return false
}

//CreateQuestion used to create new question
func (qts questionService) CreateQuestion(ctx context.Context, questionObj md.Question) bool {
	var respVal bool
	fmt.Println(`get here create question ` + questionObj.CourseName)
	createOBJ := modules.TblQuestion{CourseName: strings.ToUpper(questionObj.CourseName), Question: strings.ToUpper(questionObj.Question), CorrectOption: strings.ToUpper(questionObj.CorrectOption), CourseCode: strings.ToUpper(questionObj.CourseCode), OptionA: strings.ToUpper(questionObj.OptionA), OptionB: strings.ToUpper(questionObj.OptionB), OptionC: strings.ToUpper(questionObj.OptionC), OptionD: strings.ToUpper(questionObj.OptionD), Status: "Active", Reason: strings.ToUpper(questionObj.Reason), IsAvailable: true}
	dbGorm.Table(`tbl_question`).Create(&createOBJ)
	if createOBJ.ID > 0 {
		respVal = true
	} else {
		respVal = false
	}
	return respVal
}

//GetQuestion used to fetch questions
func (qts questionService) GetQuestion(ctx context.Context, courseName md.Question) []*md.Questions {
	//var totalCount = 0
	recsliceAll := []*md.Questions{}
	querys := dbGorm.Table(`tbl_Question`).Where(`course_name=? and is_available='false'`).First(&recsliceAll).Error
	if querys == nil {
		fmt.Println("read s")
		return recsliceAll
	}
	// query := fmt.Sprintf(`select id, course_name,question, correct_option, option_a, option_b, option_c, option_d, status, date_created, Reason from tbl_Question where course_name ='%s' and is_available = 'False'`, courseName.CourseName)
	// qoInsert, err := db.Query(query, courseName.CourseName)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }
	// fmt.Println(courseName.CourseName)
	// //qoInsert.Close()
	// for qoInsert.Next() {
	// 	quest := new(md.Questions)
	// 	totalCount++
	// 	errgg := qoInsert.Scan(&quest.ID, &quest.CourseName, &quest.Question, &quest.CorrectOption, &quest.OptionA, &quest.OptionB, &quest.OptionC, &quest.OptionD, &quest.Status, &quest.DateCreated, &quest.Reason)
	// 	if errgg == nil {
	// 		fmt.Println(quest.CourseName)
	// 		loc := new(md.Questions)
	// 		loc.ID = quest.ID
	// 		loc.CourseName = quest.CourseName
	// 		loc.Question = quest.Question
	// 		loc.CorrectOption = quest.CorrectOption
	// 		loc.OptionA = quest.OptionA
	// 		loc.OptionB = quest.OptionB
	// 		loc.OptionC = quest.OptionC
	// 		loc.OptionD = quest.OptionD
	// 		loc.Status = quest.Status
	// 		loc.DateCreated = quest.DateCreated
	// 		loc.Reason = quest.Reason
	// 		recsliceAll = append(recsliceAll, loc)
	// 	} else {
	// 		fmt.Printf(errgg.Error())
	// 	}
	//}
	return recsliceAll
}

//SubmitTestResult call to save test result GORM
func (qts questionService) SubmitTestResult(ctx context.Context, ts md.TestResult) int {
	obj := md.TblTestNewResult{StudentID: ts.StudentID, TestID: ts.TestID, TestResult: ts.TestResult, DateTaken: time.Now()}
	qts.dbGorm.Create(&obj)
	return obj.ID
}
