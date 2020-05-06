package controller

import (
	"fmt"
	"net/http"

	cr "github.com/Adebusy/dataScienceAPI/crudal"
	"github.com/Adebusy/dataScienceAPI/model"
	ut "github.com/Adebusy/dataScienceAPI/utilities"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/swag/example/celler/httputil"
)

// CreateUser godoc
// @Summary creates new user
// @Produce json
// @Param user body model.Student true "Create new user"
// @Success 200 {object} utilities.ResponseManager
// @Router /user/CreateUser [POST]
func CreateUser(c *gin.Context) {
	var userRequest model.Student
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		//check request body parameter
		fmt.Println(userRequest.Gender)
		validateRequestResp := ut.ValidateStudentReq(userRequest)
		if validateRequestResp.ResponseCode != "" {
			c.JSON(http.StatusBadRequest, validateRequestResp)
			return
		}
		doInsert := cr.CreateNewUser(userRequest)
		c.JSON(http.StatusOK, doInsert)
	}
}

// GetUserFullInfo godoc
// @Summary get user full information
// @Produce json
// @Param EmailAddress path string true "user emailAddress"
// @Success 200 {object} model.Students
// @Router /user/GetUserFullInfo/{EmailAddress} [get]
func GetUserFullInfo(ctx *gin.Context) {
	requestID := ctx.Param("EmailAddress")
	var err error
	var Msg Message
	if requestID == "" {
		Msg.Message = "Emaill address is required"
		ctx.JSON(http.StatusBadRequest, Msg)
		return
	}

	if !ut.ValidateEmail(requestID) {
		Msg.Message = "Email must be valid" + requestID
		ctx.JSON(http.StatusBadRequest, Msg)
		return
	}
	Student, err := cr.GetStudentByEmailAddress(requestID)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, Student)
}

// UpdateUserDetail godoc
// @Summary update user details
// @Produce json
// @Param user body model.Student true "Update user information"
// @Success 200 {object} controller.Message
// @Router /user/UpdateUserDetail [post]
func UpdateUserDetail(c *gin.Context) {
	var stRequest model.Student
	if err := c.ShouldBindJSON(&stRequest); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	validateRequestResp := ut.ValidateStudentReq(stRequest)
	if validateRequestResp.ResponseCode != "" {
		c.JSON(http.StatusBadRequest, validateRequestResp)
		return
	}
	//do update
	respUpdate := cr.UpdateUser(stRequest)
	c.JSON(http.StatusOK, respUpdate)
}
