package model

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
	ID          int    `json:"int" validate:"omitempty"`
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

//TestOfKnowledge student pass mark object
type TestOfKnowledge struct {
	StudentID       string `json:"StudentID"`
	Course          string `json:"MiddleName"`
	DateExamWritten string `json:"DateExamWritten"`
	Score           string `json:"Score"`
	KnowledgeLevel  string `json:"KnowledgeLevel"` //professional, entry and intermediate level
}
