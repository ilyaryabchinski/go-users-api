package users

import (
	"strconv"

	"ilyaryabchinski/gotask/src/constants"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	data, err := GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"message": constants.ErrServerError.Error(),
		})
		return
	}

	c.JSON(200, data)
}

func GetUser(c *gin.Context) {
	personalCode, parseErr := strconv.ParseUint(c.Param("personalCode"), 10, 64)

	if parseErr != nil {
		c.JSON(500, gin.H{
			"message": constants.ErrServerError.Error(),
		})
		return
	}
	user, getErr := GetOne(personalCode)
	if getErr != nil {
		c.JSON(404, gin.H{
			"message": constants.ErrNotFound.Error(),
		})
		return
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user User
	bindError := c.BindJSON(&user)

	if bindError != nil {
		c.JSON(500, gin.H{
			"message": constants.ErrServerError.Error(),
		})
		return
	}

	if user.FirstName == "" || user.LastName == "" {
		c.JSON(400, gin.H{
			"message": constants.ErrBadRequest.Error(),
		})
		return
	}

	createErr := Create(user)
	if createErr != nil {
		c.JSON(500, gin.H{
			"message": constants.ErrServerError.Error(),
		})
		return
	}
	c.JSON(200, user)

}

func UpdateUser(c *gin.Context) {
	personalCode, parseErr := strconv.ParseUint(c.Param("personalCode"), 10, 64)

	if parseErr != nil {
		c.JSON(500, gin.H{
			"message": constants.ErrServerError.Error(),
		})
		return
	}

	var user User
	bindError := c.BindJSON(&user)
	if bindError != nil {
		c.JSON(500, gin.H{
			"message": constants.ErrServerError,
		})
		return
	}

	if user.FirstName == "" || user.LastName == "" {
		c.JSON(400, gin.H{
			"message": constants.ErrBadRequest.Error(),
		})
		return
	}

	editError := Edit(personalCode, user)
	if editError != nil {
		c.JSON(404, gin.H{
			"message": constants.ErrNotFound.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	personalCode, parseErr := strconv.ParseUint(c.Param("personalCode"), 10, 64)

	if parseErr != nil {
		c.JSON(500, gin.H{
			"message": constants.ErrServerError.Error(),
		})
		return
	}
	deleteError := Delete(personalCode)
	if deleteError != nil {
		c.JSON(404, gin.H{
			"message": constants.ErrNotFound.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": constants.DeleteSuccess,
	})
}
