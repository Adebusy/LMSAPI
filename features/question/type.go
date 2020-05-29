package question

import (
	"regexp"
	"time"

	"github.com/Adebusy/dataScienceAPI/modules"
)

//QuisRequest response
type QuisRequest struct {
	MyQuestion     []*modules.Questions
	StudentDetails modules.Students
	CourseDetails  modules.TblCourse
}

//TblTestNewResult of students
type TblTestNewResult struct {
	ID         int `gorm:"primary_key;auto_increment:true"`
	StudentID  string
	TestID     string
	TestResult string
	DateTaken  time.Time //time.Time
}

//CREATE TABLE "tbl_test_result" ("ID" INTEGER,"student_id" VARCHAR(255),"test_id" VARCHAR(255),"test_result" VARCHAR(255),"date_taken" VARCHAR(255) )
//	ID              int    `json:"ID" validate:"omitempty"`
//RequestResponse response
type RequestResponse struct {
	ResponseMessage string `json:"ResponseMessage" validate:"omitempty"`
	ResponseCode    string `json:"ResponseCode" validate:"omitempty"`
}

//GetQuestionObj obj for GetQuestion
type GetQuestionObj struct {
	CourseName    string `json:"CourseName" validate:"omitempty"`
	QuestionCount int    `json:"QuestionCount" validate:"omitempty"`
}

//ValidateQuestionReq check request
func ValidateQuestionReq(questObj modules.Question) modules.RequestResponse {
	var resp modules.RequestResponse
	resp.ResponseCode = ""
	resp.ResponseMessage = ""
	if questObj.CorrectOption == "" {
		resp.ResponseMessage = "Correction option is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.CourseName == "" {
		resp.ResponseMessage = "CourseName is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.OptionA == "" {
		resp.ResponseMessage = "Atleast option A is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.OptionB == "" {
		resp.ResponseMessage = "Atleast option B is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.Reason == "" {
		resp.ResponseMessage = "Reason is required"
		resp.ResponseCode = "01"
		return resp
	}

	return resp
}

//ValidateEmail used to validate email address
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}
