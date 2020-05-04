package utilities

import (
	"log"
	"os"
	"regexp"

	"github.com/Adebusy/dataScienceAPI/model"
	"github.com/joho/godotenv"
)

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

//ValidateStudentReq used to validate office request
func ValidateStudentReq(student model.Student) ResponseManager {
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

//ValidateQuestionReq used to validate question request
func ValidateQuestionReq(questObj model.Question) ResponseManager {
	var resp ResponseManager
	resp.ResponseCode = ""
	resp.ResponseDescription = ""
	if questObj.CorrectOption == "" {
		resp.ResponseDescription = "Correction option is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.CourseName == "" {
		resp.ResponseDescription = "CourseName is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.OptionA == "" {
		resp.ResponseDescription = "Atleast option A is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.OptionB == "" {
		resp.ResponseDescription = "Atleast option B is required"
		resp.ResponseCode = "01"
		return resp
	}

	if questObj.Reason == "" {
		resp.ResponseDescription = "Reason is required"
		resp.ResponseCode = "01"
		return resp
	}

	return resp
}

//ResponseManager general response
type ResponseManager struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"`
	ResponseCode        string `json:"ResponseCode,omitempty"`
}

//Checkerr checking error
func Checkerr(err error) {
	if err != nil {
		panic(err.Error)
	}
}
