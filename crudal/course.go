package crudal

import (
	"fmt"
	"log"
	"strings"

	"github.com/Adebusy/dataScienceAPI/model"
)

//234012803823  08080523171

//CheckIfCourseExist validates course name
func CheckIfCourseExist(courseNam string) model.Courses {
	var courseobject model.Courses
	query := fmt.Sprintf(`select * from tbl_Courses where CourseName ='%s'`, courseNam)
	//fmt.Println(query)
	doinsert, err := db.Query(query, courseNam)
	if err != nil {
		log.Panic(err.Error())
	}
	if doinsert.Next() == true {
		for doinsert.Next() {
			respvals := doinsert.Scan(&courseobject.ID, &courseobject.CourseName, &courseobject.CourseCode, &courseobject.CourseCategory, &courseobject.CourseStatus, &courseobject.DateCreated, &courseobject.QuestionCount, &courseobject.PassMark)
			if respvals != nil {
				log.Printf(respvals.Error())
			}
		}
	} else {

	}
	return courseobject
}

//CheckIfCourseExistBool validates course name and return true false
func CheckIfCourseExistBool(courseNam string) bool {
	query := fmt.Sprintf(`select * from tbl_Courses where CourseName ='%s'`, courseNam)
	//fmt.Println(query)
	doinsert, err := db.Query(query, courseNam)
	if err != nil {
		log.Panic(err.Error())
	}
	if doinsert.Next() == true {
		//fmt.Println("seen a record CheckIfCourseExistBool")
		return true
	}
	return false
}

var respVal bool

//CreateNewCourse create a new course
func CreateNewCourse(courseOBJ model.Course, getcourseCode string) bool {
	respVal = false
	var newQury = fmt.Sprintf(`insert into tbl_Courses(CourseName, CourseCode, CourseCategory, QuestionCount, passMark)values ('%s','%s','%s','%s','%s')`, strings.ToUpper(courseOBJ.CourseName), strings.ToUpper(getcourseCode), courseOBJ.CourseCategory, courseOBJ.QuetionCount, courseOBJ.PassMark)
	//fmt.Println(newQury)
	query, _ := db.Prepare(newQury)
	_, err := query.Exec()
	if err != nil {
		fmt.Printf(err.Error())
		respVal = false
	} else {
		respVal = true
	}
	defer db.Close()
	return respVal
}
