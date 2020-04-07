package user

import (
	"bookstore/domain/user"
	usrSrvc "bookstore/services"
	appErr "bookstore/utils/error"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userInfo user.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		appErr := appErr.AppErr{Message: "Unable to read body", Status: http.StatusBadRequest, Error: "bad_request"}
		c.JSON(appErr.Status, appErr)
		return
	}
	if err = json.Unmarshal(bytes, &userInfo); err != nil {
		appErr := appErr.AppErr{Message: "Invalid JSON body", Status: http.StatusBadRequest, Error: "bad_request"}
		c.JSON(appErr.Status, appErr)
		return
	}
	fmt.Println(userInfo)

	createdUser, appErr := usrSrvc.UserService.CreateUser(userInfo)
	if appErr != nil {
		c.JSON(appErr.Status, appErr)
		return
	}

	c.JSON(http.StatusCreated, createdUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {
	userId, usrErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if usrErr != nil {
		err := appErr.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, appErr := usrSrvc.UserService.GetUser(userId)
	if appErr != nil {
		c.JSON(appErr.Status, appErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	var userInfo user.User
	userId, usrErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if usrErr != nil {
		err := appErr.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	// short way to bind JSON body with Struct
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		appErr := appErr.NewBadRequestError("Invalid JSON body")
		c.JSON(appErr.Status, appErr)
		return
	}

	userInfo.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	updatedUser, appErr := usrSrvc.UserService.UpdateUser(isPartial, userInfo)
	if appErr != nil {
		c.JSON(appErr.Status, appErr)
		return
	}
	c.JSON(http.StatusOK, updatedUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		err := appErr.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	appErr := usrSrvc.UserService.DeleteUser(userId)
	if appErr != nil {
		c.JSON(appErr.Status, appErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")

	users, appErr := usrSrvc.UserService.FindByStatus(status)
	if appErr != nil {
		c.JSON(appErr.Status, appErr)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
