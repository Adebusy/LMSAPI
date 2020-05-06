package crudal

import (
	"fmt"
	"log"
	"strings"

	"github.com/Adebusy/dataScienceAPI/model"
	utl "github.com/Adebusy/dataScienceAPI/utilities"
)

//GetStudentByEmailAddress read student detail from db
func GetStudentByEmailAddress(EmaillAddress string) (model.Students, error) {
	var userOBJ model.Students
	query := fmt.Sprintf(`select ID,FirstName,MiddleName,LastName,Gender ,Email,Nationality, Category,Status ,DateCreated ,UserType from tbl_Students where Email ='%s'`, EmaillAddress)
	resp, err := db.Query(query, EmaillAddress)
	resp.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		for resp.Next() {
			err := resp.Scan(&userOBJ.ID, &userOBJ.FirstName, &userOBJ.MiddleName, &userOBJ.LastName, &userOBJ.Gender, &userOBJ.Email, &userOBJ.Nationality, &userOBJ.Category, &userOBJ.Status, &userOBJ.DateCreated, &userOBJ.UserType)
			if err != nil {
				log.Print(err)
				return userOBJ, err
			}
			return userOBJ, err
		}
	}
	return userOBJ, err
}

//CreateNewUser method to create new user
func CreateNewUser(RequesuetOBJ model.Student) utl.ResponseManager {
	var resp utl.ResponseManager
	query := fmt.Sprintf(`insert into tbl_Students (FirstName,MiddleName,LastName,Gender,Email,Nationality,Category) values ('%s','%s','%s','%s','%s','%s','%s')`, strings.ToUpper(RequesuetOBJ.FirstName), strings.ToUpper(RequesuetOBJ.MiddleName), strings.ToUpper(RequesuetOBJ.LastName), strings.ToUpper(RequesuetOBJ.Gender), strings.ToUpper(RequesuetOBJ.Email), strings.ToUpper(RequesuetOBJ.Nationality), strings.ToUpper(RequesuetOBJ.Category))
	respFromCheck, errcheck := GetStudentByEmailAddress(strings.ToUpper(RequesuetOBJ.Email))
	if errcheck != nil {
		//log.Panic(errcheck.Error)
		resp.ResponseCode = "01"
		resp.ResponseDescription = "Unable to validate Email address at the moment. Please try again later."
		return resp
	}
	if respFromCheck.FirstName != "" {
		resp.ResponseCode = "01"
		resp.ResponseDescription = "A user with Email address supplied already exist."
		return resp
	}
	doRequest, err := db.Prepare(query)
	if err != nil {
		log.Panic(err)
		resp.ResponseCode = "01"
		resp.ResponseDescription = err.Error()
		return resp
	}
	_, errs := doRequest.Exec()
	if errs != nil {
		log.Panic(errs)
		resp.ResponseCode = "01"
		resp.ResponseDescription = errs.Error()
	} else {
		resp.ResponseCode = "00"
		resp.ResponseDescription = "User created successfully"
	}

	return resp
}

//UpdateUser method to create new user
func UpdateUser(RequesuetOBJ model.Student) utl.ResponseManager {
	var resp utl.ResponseManager
	query := fmt.Sprintf(`update tbl_Students  set FirstName='%s', MiddleName='%s', LastName='%s', Gender='%s', Nationality='%s', Category='%s', UserType='%s' where Email='%s'`, RequesuetOBJ.FirstName, RequesuetOBJ.MiddleName, RequesuetOBJ.LastName, RequesuetOBJ.Gender, RequesuetOBJ.Nationality, RequesuetOBJ.Category, RequesuetOBJ.UserType, RequesuetOBJ.Email)
	//check user exist
	respFromCheck, errcheck := GetStudentByEmailAddress(RequesuetOBJ.Email)
	if errcheck != nil {
		log.Panic(errcheck.Error)
		resp.ResponseCode = "01"
		resp.ResponseDescription = "Unable to validate Email address at the moment. Please try again later."
		return resp
	}
	if respFromCheck.FirstName == "" {
		resp.ResponseCode = "01"
		resp.ResponseDescription = "Email address does not exist."
	}

	doRequest, err := db.Prepare(query)
	if err != nil {
		log.Panic(err)
		resp.ResponseCode = "01"
		resp.ResponseDescription = err.Error()
		return resp
	}
	_, errs := doRequest.Exec()
	if errs != nil {
		log.Panic(errs)
		resp.ResponseCode = "01"
		resp.ResponseDescription = errs.Error()
	} else {
		resp.ResponseCode = "00"
		resp.ResponseDescription = "User updated successfully."
	}
	return resp
}
