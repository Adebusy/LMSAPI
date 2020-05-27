package student

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	st "github.com/Adebusy/dataScienceAPI/datastore/sqlserver/student"
	"github.com/Adebusy/dataScienceAPI/driver/sqlserver"
	md "github.com/Adebusy/dataScienceAPI/modules"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
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

var Istudent = st.NewstudentService(db)
var resp md.ResponseManager

// CreateUser godoc
// @Summary creates new user
// @Produce json
// @Param user body modules.Student true "Create new user"
// @Success 200 {object} modules.ResponseManager
// @Router /user/CreateUser [POST]
func CreateUser(c *gin.Context) {
	var userRequest md.TblStudent
	//var resp md.ResponseManager
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		//check request body parameter
		fmt.Println(userRequest.Gender)
		validateRequestResp := md.ValidateStudentReq(userRequest)
		if validateRequestResp.ResponseCode != "" {
			c.JSON(http.StatusBadRequest, validateRequestResp)
			return
		}
		respFromCheck, errcheck := Istudent.GetStudentByEmailAddress(c, userRequest.Email)
		if errcheck != nil {
			resp.ResponseCode = "01"
			resp.ResponseDescription = "Unable to validate Email address at the moment. Please try again later."
			// c.JSON(http.StatusOK, errcheck)
			// return
			fmt.Println("get hgere")
			fmt.Println(errcheck)
		}

		if respFromCheck.FirstName != "" {
			resp.ResponseCode = "01"
			resp.ResponseDescription = "A user with Email address supplied already exist."
			c.JSON(http.StatusOK, resp)
			return
		}
		doInsert := Istudent.Create(c, userRequest)
		c.JSON(http.StatusOK, doInsert)
	}
}

// GetUserFullInfo godoc
// @Summary get user full information
// @Produce json
// @Param EmailAddress path string true "user emailAddress"
// @Success 200 {object} modules.Students
// @Router /user/GetUserFullInfo/{EmailAddress} [get]
func GetUserFullInfo(ctx *gin.Context) {
	requestID := ctx.Param("EmailAddress")
	var err error
	if requestID == "" {
		resp.ResponseDescription = "Emaill address is required"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if !md.ValidateEmail(requestID) {
		resp.ResponseDescription = "Email must be valid" + requestID
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	Student, err := Istudent.GetStudentByEmailAddress(ctx, requestID)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		ctx.JSON(http.StatusBadRequest, Student)
		return
	}
	ctx.JSON(http.StatusOK, Student)
}

// UpdateUserDetail godoc
// @Summary update user details
// @Produce json
// @Param user body modules.Student true "Update user information"
// @Success 200 {object} modules.Message
// @Router /user/UpdateUserDetail [post]
func UpdateUserDetail(c *gin.Context) {
	var stRequest md.TblStudent
	if err := c.ShouldBindJSON(&stRequest); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	validateRequestResp := md.ValidateStudentReq(stRequest)
	if validateRequestResp.ResponseCode != "" {
		c.JSON(http.StatusBadRequest, validateRequestResp)
		return
	}

	respFromCheck, errcheck := Istudent.GetStudentByEmailAddress(c, stRequest.Email)
	if errcheck != nil {
		fmt.Println(errcheck)
	}

	if respFromCheck.FirstName == "" {
		resp.ResponseCode = "01"
		resp.ResponseDescription = "User with Email address supplied does not exist."
		c.JSON(http.StatusOK, resp)
		return
	}
	//do update
	doUpdate := Istudent.Update(c, stRequest)
	c.JSON(http.StatusOK, doUpdate)
}
