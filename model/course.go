package model

//Course entity
type Course struct {
	CourseName     string `json:"CourseName" validate:"omitempty"`
	CourseCategory string `json:"CourseCategory" validate:"omitempty"`
	QuetionCount   string `json:"QuetionCount" validate:"omitempty"`
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
