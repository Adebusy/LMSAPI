package model

import "time"

//Question list of  entity
type Question struct {
	CourseName    string `json:"CourseName" validate:"omitempty"`
	Question      string `json:"Question" validate:"omitempty"`
	CorrectOption string `json:"CorrectOption" validate:"omitempty"`
	OptionA       string `json:"OptionA" validate:"omitempty"`
	OptionB       string `json:"OptionB" validate:"omitempty"`
	OptionC       string `json:"OptionC" validate:"omitempty"`
	OptionD       string `json:"OptionD" validate:"omitempty"`
	Reason        string `json:"Reason" validate:"omitempty"`
}

//Questions list of  entity
type Questions struct {
	ID            int    `json:"ID" validate:"omitempty"`
	CourseName    string `json:"CourseName" validate:"omitempty"`
	Question      string `json:"Question" validate:"omitempty"`
	CorrectOption string `json:"CorrectOption" validate:"omitempty"`
	CourseCode    string `json:"CourseCode" validate:"omitempty"`
	//CourseCategory string `json:"CourseCategory" validate:"omitempty"`
	OptionA     string `json:"OptionA" validate:"omitempty"`
	OptionB     string `json:"OptionB" validate:"omitempty"`
	OptionC     string `json:"OptionC" validate:"omitempty"`
	OptionD     string `json:"OptionD" validate:"omitempty"`
	Status      string `json:"Status" validate:"omitempty"` //Active , non-Active
	DateCreated string `json:"DateCreated" validate:"omitempty"`
	Reason      string `json:"Reason" validate:"omitempty"`
	IsAvailable bool   `json:"IsAvailable" validate:"omitempty"`
	RecCount    int    `json:"RecCount" validate:"omitempty"`
}

//QuisRequest response
type QuisRequest struct {
	MyQuestion     []*Questions
	StudentDetails Students
	CourseDetails  Courses
}

//TestResult of students
type TestResult struct {
	StudentID  string `json:"StudentID" validate:"omitempty"`
	TestID     string `json:"TestID" validate:"omitempty"`
	TestResult string `json:"TestResult" validate:"omitempty"`
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
