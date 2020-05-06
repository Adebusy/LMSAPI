package crudal

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/Adebusy/dataScienceAPI/model"
)

var recsliceRaw = []*model.Questions{}
var randonNumbers int = 0

//CreateQuestion used to create new question
func CreateQuestion(questionObj model.Question) bool {
	var respVal bool
	query := fmt.Sprintf(`insert into tbl_Questions(CourseName,Question,CorrectOption,OptionA,OptionB,OptionC,OptionD,Reason) values ('%s','%s','%s','%s','%s','%s','%s','%s')`, strings.ToUpper(questionObj.CourseName), questionObj.Question, questionObj.CorrectOption, questionObj.OptionA, questionObj.OptionB, questionObj.OptionC, questionObj.OptionD, questionObj.Reason)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf(err.Error())
		log.Panic(err.Error)
		respVal = false
		return respVal
	}
	respVal = true
	return respVal
}

//GetQuestion used to fetch questions
func GetQuestion(courseName string, QuestionCount int) []*model.Questions {
	var totalCount = 0
	recsliceAll := []*model.Questions{}
	query := fmt.Sprintf(`select ID, CourseName,Question, CorrectOption, OptionA, OptionB, OptionC, OptionD, Status, DateCreated, Reason from tbl_Questions where CourseName ='%s' and isAvailable = 1`, courseName)
	//fmt.Println(query)
	qoInsert, err := db.Query(query, courseName)
	if err != nil {
		fmt.Printf(err.Error())
	}
	qoInsert.Close()
	for qoInsert.Next() {
		quest := new(model.Questions)
		totalCount++
		errgg := qoInsert.Scan(&quest.ID, &quest.CourseName, &quest.Question, &quest.CorrectOption, &quest.OptionA, &quest.OptionB, &quest.OptionC, &quest.OptionD, &quest.Status, &quest.DateCreated, &quest.Reason)
		if errgg == nil {
			fmt.Println(quest.CourseName)
			loc := new(model.Questions)
			loc.ID = quest.ID
			loc.CourseName = quest.CourseName
			loc.Question = quest.Question
			loc.CorrectOption = quest.CorrectOption
			loc.OptionA = quest.OptionA
			loc.OptionB = quest.OptionB
			loc.OptionC = quest.OptionC
			loc.OptionD = quest.OptionD
			loc.Status = quest.Status
			loc.DateCreated = quest.DateCreated
			loc.Reason = quest.Reason
			recsliceAll = append(recsliceAll, loc)
		} else {
			fmt.Printf(errgg.Error())
		}
		// for i := 1; i <= QuestionCount; i++ {
		// 	//generate random number with range
		// 	for (checkAlreadyAdd(recsliceRaw, totalCount)) == false {
		// 		doAdd(recslice)
		// 		fmt.Println("get here p2")
		// 		break
		// 	}
		// }
	}
	return recsliceAll
}
func doAdd(recslice []*model.Questions) {
	for _, value := range recslice {
		if value.RecCount == randonNumbers {
			recsliceRaw = append(recsliceRaw, value)
			break
		}
	}
}

func checkAlreadyAdd(recsliceRaw []*model.Questions, randomMax int) bool {
	var retCheck = false
	randonNumbers = rand.Intn(randomMax)
	//check if it has'nt been added before
	for _, rawCheck := range recsliceRaw {
		if rawCheck.RecCount == randonNumbers {
			retCheck = true
			//fmt.Println("record with id " + strconv.Itoa(randonNumbers))
			break
		}
	}
	return retCheck
}
