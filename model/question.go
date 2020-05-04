package model

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
	OptionA       string `json:"OptionA" validate:"omitempty"`
	OptionB       string `json:"OptionB" validate:"omitempty"`
	OptionC       string `json:"OptionC" validate:"omitempty"`
	OptionD       string `json:"OptionD" validate:"omitempty"`
	Status        string `json:"Status" validate:"omitempty"` //Active , non-Active
	DateCreated   string `json:"DateCreated" validate:"omitempty"`
	Reason        string `json:"Reason" validate:"omitempty"`
	IsAvailable   bool   `json:"IsAvailable" validate:"omitempty"`
	RecCount      int    `json:"RecCount" validate:"omitempty"`
}

//QuisRequest response
type QuisRequest struct {
	MyQuestion     []*Questions
	StudentDetails Students
	CourseDetails  Courses
}
