package student

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Adebusy/dataScienceAPI/driver/sqlserver"
	md "github.com/Adebusy/dataScienceAPI/modules"
	"github.com/jinzhu/gorm"
)

var db *sql.DB
var err error

var dbGorm *gorm.DB
var errGorm error

type courseService struct {
	db     *sql.DB
	dbGorm *gorm.DB
}

func init() {
	dbGorm, errGorm = sqlserver.ConnectGorm()
	db, err = sqlserver.ConnectOLEDBC()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no connection error")
	}
	if errGorm != nil {
		fmt.Printf(errGorm.Error())
	} else {

	}

}

type StudentService interface {
	GetStudentByEmailAddress(ctx context.Context, EmaillAddress string) (md.Students, error)
	Create(ctx context.Context, RequesuetOBJ md.TblStudent) md.ResponseManager
	Update(ctx context.Context, RequesuetOBJ md.TblStudent) md.ResponseManager
}

//NewstudentService interface to be assesible
func NewstudentService(db *sql.DB) StudentService {
	return &studentService{db, dbGorm}
}

type studentService struct {
	db     *sql.DB
	dbGorm *gorm.DB
}

//GetStudentByEmailAddress read student detail from db
func (ts studentService) GetStudentByEmailAddress(ctx context.Context, EmaillAddress string) (md.Students, error) {
	var userOBJ md.Students
	errquery := dbGorm.Table("tbl_student").Where(`Email=?`, EmaillAddress).First(&userOBJ).Error
	if errquery == nil {
		return userOBJ, nil
	}
	return userOBJ, errquery
}

//Create method to create new user
func (ts studentService) Create(ctx context.Context, RequesuetOBJ md.TblStudent) md.ResponseManager {
	var resp md.ResponseManager
	RequesuetOBJ.DateCreated = time.Now()
	RequesuetOBJ.Status = "1"
	RequesuetOBJ.Category = "Entry Level"
	query := dbGorm.Table("tbl_student").Create(&RequesuetOBJ)
	if query == nil {
		resp.ResponseCode = "01"
		resp.ResponseDescription = query.Error.Error()
		fmt.Printf("created request for new student with email address = %s", strconv.Itoa(RequesuetOBJ.ID))
	} else {
		resp.ResponseCode = "00"
		resp.ResponseDescription = "User created successfully"
	}
	return resp
}

//Update method to create new user
func (ts studentService) Update(ctx context.Context, RequesuetOBJ md.TblStudent) md.ResponseManager {
	var resp md.ResponseManager
	query := fmt.Sprintf(`update tbl_Student  set first_name='%s', middle_name='%s', last_name='%s', gender='%s', nationality='%s', category='%s', user_type='%s' where email='%s'`, RequesuetOBJ.FirstName, RequesuetOBJ.MiddleName, RequesuetOBJ.LastName, RequesuetOBJ.Gender, RequesuetOBJ.Nationality, RequesuetOBJ.Category, RequesuetOBJ.UserType, RequesuetOBJ.Email)
	doRequest, _ := db.Prepare(query)
	_, errs := doRequest.Exec()
	if errs != nil {
		resp.ResponseCode = "01"
		resp.ResponseDescription = errs.Error()
	} else {
		resp.ResponseCode = "00"
		resp.ResponseDescription = "updated successfully."
	}
	return resp
}
