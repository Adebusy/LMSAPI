package modules

import "time"

//Course list of  entity
type Course struct {
	CourseName     string `json:"CourseName" validate:"omitempty"`
	CourseCode     string `json:"CourseCode" validate:"omitempty"`
	CourseCategory string `json:"CourseCategory" validate:"omitempty"`
	CourseStatus   string `json:"CourseStatus" validate:"omitempty"`
	DateCreated    string `json:"DateCreated" validate:"omitempty"`
	QuestionCount  string `json:"QuestionCount" validate:"omitempty"`
	PassMark       string `json:"PassMark" validate:"omitempty"`
}

//Courses list of  entity
type Courses struct {
	ID             int    `json:"ID" validate:"omitempty"`
	CourseName     string `json:"CourseName" validate:"omitempty"`
	CourseCode     string `json:"CourseCode" validate:"omitempty"`
	CourseCategory string `json:"CourseCategory" validate:"omitempty"`
	CourseStatus   string `json:"CourseStatus" validate:"omitempty"`
	DateCreated    string `json:"DateCreated" validate:"omitempty"`
	QuestionCount  string `json:"QuestionCount" validate:"omitempty"`
	PassMark       string `json:"PassMark" validate:"omitempty"`
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}

//CreateObj  for obj
type CreateObj struct {
	c             Course
	getcourseCode string
}

//TblCourse list of  entity
type TblCourse struct {
	ID             int       `json:"ID" validate:"omitempty"`
	CourseName     string    `json:"CourseName" validate:"omitempty"`
	CourseCode     string    `json:"CourseCode" validate:"omitempty"`
	CourseCategory string    `json:"CourseCategory" validate:"omitempty"`
	CourseStatus   string    `json:"CourseStatus" validate:"omitempty"`
	DateCreated    time.Time `json:"DateCreated" validate:"omitempty"`
	QuestionCount  string    `json:"QuestionCount" validate:"omitempty"`
	PassMark       string    `json:"PassMark" validate:"omitempty"`
}
