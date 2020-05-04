package crudal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/Adebusy/VisitorsManager/AppCode"
	"github.com/Adebusy/dataScienceAPI/model"
	"github.com/joho/godotenv"
)

var recsliceRaw = []*model.Questions{}
var randonNumbers int = 0

func init() {
	godotenv.Load()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		AppCode.GoDotEnvVariable("Server"), AppCode.GoDotEnvVariable("user"), AppCode.GoDotEnvVariable("Password"), AppCode.GoDotEnvVariable("Port"), AppCode.GoDotEnvVariable("Database"))
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	} else {
		fmt.Println("no eerror")
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no eerror 2")
	}
}

//CreateQuestion used to create new question
func CreateQuestion(questionObj model.Question) bool {
	var respVal bool
	fmt.Println("got here 1")
	query := fmt.Sprintf(`insert into tbl_Questions(CourseName,Question,CorrectOption,OptionA,OptionB,OptionC,OptionD,Reason) values ('%s','%s','%s','%s','%s','%s','%s','%s')`, strings.ToUpper(questionObj.CourseName), questionObj.Question, questionObj.OptionA, questionObj.OptionB, questionObj.OptionC, questionObj.OptionD, questionObj.Reason)
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
	recslice := []*model.Questions{}
	recsliceRaw := []*model.Questions{}
	var quest model.Questions
	query := fmt.Sprintf(`select * from tbl_Questions where CourseName ='%s' order by ID `, courseName)
	if checkdbstatus := db.Ping == nil; checkdbstatus == true {
		qoInsert, err := db.Query(query, courseName)
		if err != nil {
			log.Panic(err.Error)
		} else {
			for qoInsert.Next() {
				totalCount++
				qoInsert.Scan(&quest)
				loc := new(model.Questions)
				loc.RecCount = totalCount
				loc.ID = quest.ID
				loc.CorrectOption = quest.CorrectOption
				loc.CourseName = quest.CourseName
				loc.DateCreated = quest.DateCreated
				loc.OptionA = quest.OptionA
				loc.OptionB = quest.OptionB
				loc.OptionC = quest.OptionC
				loc.OptionD = quest.OptionD
				loc.Question = quest.Question
				loc.Reason = quest.Reason
				loc.Status = quest.Status
				recslice = append(recslice, loc)
			}
			for i := 1; i <= QuestionCount; i++ {
				//generate random number with range
				for (checkAlreadyAdd(recsliceRaw, totalCount)) == false {
					doAdd(recslice)
					break
				}
			}
		}
	}
	return recsliceRaw
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
			break
		}
	}
	return retCheck
}
