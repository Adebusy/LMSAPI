package modules

import "time"

//RequestResponse response
type RequestResponse struct {
	ResponseMessage string `json:"ResponseMessage" validate:"omitempty"`
	ResponseCode    string `json:"ResponseCode" validate:"omitempty"`
}

//TestResult object
type TestResult struct {
	StudentID  string `json:"StudentID"`
	TestID     string `json:"TestID"`
	TestResult string `json:"TestResult"`
}

//Question list of  entity
type Question struct {
	CourseName    string `json:"CourseName" validate:"omitempty"`
	Question      string `json:"Question" validate:"omitempty"`
	CorrectOption string `json:"CorrectOption" validate:"omitempty"`
	CourseCode    string `json:"CourseCode" validate:"omitempty"`
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

//TblTestNewResult for gorm ORM
type TblTestNewResult struct {
	ID         int       `json:"ID" validate:"omitempty"`
	StudentID  string    `json:"StudentID" validate:"omitempty"`
	TestID     string    `json:"TestID" validate:"omitempty"`
	TestResult string    `json:"TestResult" validate:"omitempty"`
	DateTaken  time.Time `json:"DateTaken" validate:"omitempty"`
}

type TblQuestion struct {
	ID            int    `json:"ID" validate:"omitempty"`
	CourseName    string `json:"CourseName" validate:"omitempty"`
	Question      string `json:"Question" validate:"omitempty"`
	CorrectOption string `json:"CorrectOption" validate:"omitempty"`
	CourseCode    string `json:"CourseCode" validate:"omitempty"`
	//CourseCategory string `json:"CourseCategory" validate:"omitempty"`
	OptionA     string    `json:"OptionA" validate:"omitempty"`
	OptionB     string    `json:"OptionB" validate:"omitempty"`
	OptionC     string    `json:"OptionC" validate:"omitempty"`
	OptionD     string    `json:"OptionD" validate:"omitempty"`
	Status      string    `json:"Status" validate:"omitempty"` //Active , non-Active
	DateCreated time.Time `json:"DateCreated" validate:"omitempty"`
	Reason      string    `json:"Reason" validate:"omitempty"`
	IsAvailable bool      `json:"IsAvailable" validate:"omitempty"`
	RecCount    int       `json:"RecCount" validate:"omitempty"`
}
