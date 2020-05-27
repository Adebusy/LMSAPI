package modules

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/joho/godotenv"
)

//Student student entity
type Student struct {
	FirstName   string `json:"FirstName" validate:"omitempty"`
	MiddleName  string `json:"MiddleName" validate:"omitempty"`
	LastName    string `json:"LastName" validate:"omitempty"`
	Gender      string `json:"Gender" validate:"omitempty"`
	Email       string `json:"Email" validate:"omitempty"`
	Nationality string `json:"Nationality" validate:"omitempty"`
	Category    string `json:"Category" validate:"omitempty"`
	UserType    string `json:"UserType" validate:"omitempty"`
}

//Students student entity
type Students struct {
	ID          int    `json:"ID" validate:"omitempty"`
	FirstName   string `json:"FirstName" validate:"omitempty"`
	MiddleName  string `json:"MiddleName" validate:"omitempty"`
	LastName    string `json:"LastName" validate:"omitempty"`
	Gender      string `json:"Gender" validate:"omitempty"`
	Email       string `json:"Email" validate:"omitempty"`
	Nationality string `json:"Nationality" validate:"omitempty"`
	Category    string `json:"Category" validate:"omitempty"`
	Status      string `json:"Status" validate:"omitempty"`
	DateCreated string `json:"DateCreated" validate:"onitempty"`
	UserType    string `json:"UserType" validate:"onitempty"`
}

//ResponseManager general response
type ResponseManager struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"`
	ResponseCode        string `json:"ResponseCode,omitempty"`
}

//TestOfKnowledge student pass mark object
type TestOfKnowledge struct {
	StudentID       string `json:"StudentID"`
	Course          string `json:"MiddleName"`
	DateExamWritten string `json:"DateExamWritten"`
	Score           string `json:"Score"`
	KnowledgeLevel  string `json:"KnowledgeLevel"` //professional, entry and intermediate level
}

//ValidateStudentReq used to validate office request
func ValidateStudentReq(student TblStudent) ResponseManager {
	var resp ResponseManager
	resp.ResponseCode = ""
	resp.ResponseDescription = ""
	if student.FirstName == "" {
		resp.ResponseDescription = "Firstname is required"
		resp.ResponseCode = "01"
		return resp
	}

	if student.Email == "" {
		resp.ResponseDescription = "Email is required"
		resp.ResponseCode = "01"
		return resp
	}
	if !ValidateEmail(student.Email) {
		resp.ResponseDescription = "Invalid Email supplied"
		resp.ResponseCode = "01"
		return resp
	}

	if student.Gender == "" {
		resp.ResponseDescription = "Gender is required"
		resp.ResponseCode = "01"
		return resp
	}

	if student.Gender != "F" && student.Gender != "M" {
		resp.ResponseDescription = "Gender must be F for female and M for male"
		resp.ResponseCode = "01"
		return resp
	}

	if student.LastName == "" {
		resp.ResponseDescription = "Lastname is required"
		resp.ResponseCode = "01"
		return resp
	}

	if student.Nationality == "" {
		resp.ResponseDescription = "Nationality is required"
		resp.ResponseCode = "01"
		return resp
	}
	return resp
}

//Checkerr checking error
func Checkerr(err error) {
	if err != nil {
		fmt.Println(err.Error)
	}
}

//ValidateEmail used to validate email address
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

func getVariables(key string) string {
	os.Setenv("password", "tylent")
	os.Setenv("dbname", "VisitorDB")
	os.Setenv("server", "10.0.41.101")
	os.Setenv("databaseSchema", "mysql")
	os.Setenv("root", "root")
	os.Setenv(key, "VisitorDB")
	return os.Getenv(key)
}

//GoDotEnvVariable load env file
func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

//TblStudent student entity
type TblStudent struct {
	ID          int       `json:"int" validate:"omitempty"`
	FirstName   string    `json:"FirstName" validate:"omitempty"`
	MiddleName  string    `json:"MiddleName" validate:"omitempty"`
	LastName    string    `json:"LastName" validate:"omitempty"`
	Gender      string    `json:"Gender" validate:"omitempty"`
	Email       string    `json:"Email" validate:"omitempty"`
	Nationality string    `json:"Nationality" validate:"omitempty"`
	Category    string    `json:"Category" validate:"omitempty"`
	Status      string    `json:"Status" validate:"omitempty"`
	DateCreated time.Time `json:"DateCreated" validate:"onitempty"`
	UserType    string    `json:"UserType" validate:"onitempty"`
}
