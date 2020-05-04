package controller

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	crd "github.com/Adebusy/dataScienceAPI/crudal"
	"github.com/Adebusy/dataScienceAPI/model"
	"github.com/gin-gonic/gin"
)

// CreateCourse godoc
// @Summary create new course
// @Produce json
// @Param user body model.Course true "Create new course"
// @Success 200 {object} controller.Message
// @Router /course/CreateCourse [post]
func CreateCourse(ctx *gin.Context) {
	var course model.Course
	var msg Message
	ctx.ShouldBindJSON(&course)
	if len(course.CourseName) < 3 {
		msg.Message = "Course name length must be greated than 3."
		ctx.JSON(http.StatusBadRequest, msg)
	}
	//check if course name already exist
	checkAlreadyExist := crd.CheckIfCourseExist(course.CourseName)
	if checkAlreadyExist.CourseCode == "" {
		n := rand.Int63()
		var getcourseCode = strings.ToUpper(course.CourseName[0:3]) + strconv.FormatInt(n, 2)
		if respCreate := crd.CreateNewCourse(course, getcourseCode); respCreate == true {
			msg.Message = "New course with code '" + getcourseCode + "' created successfully."
			ctx.JSON(http.StatusOK, msg)
		} else {
			msg.Message = "Unable to create new course successfully"
			ctx.JSON(http.StatusOK, msg)
		}
	} else {
		msg.Message = "Name already exist."
		ctx.JSON(http.StatusOK, msg)
	}
}
