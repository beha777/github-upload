package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"testNEW/db"
	"testNEW/models"
)

func userRegistration (context *gin.Context) {

	var user models.User
	err := json.NewDecoder(context.Request.Body).Decode(&user)
	fmt.Println("-------------", user)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"response": err.Error(),
		})
		return
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"response": err.Error(),
		})
		return
	}
	user.Password = string(bytes)

	if db.IsExistUser(user.Login) {
		context.JSON(http.StatusOK, gin.H{
			"response": "User with such login already exists",
		})
	} else {
		db.GetDBConn().Create(&user)
		context.JSON(http.StatusOK, gin.H{
			"response": "You have successfully registered",
		})
	}
}


func userByName(context *gin.Context) {
	userName := context.Query("name")


	fmt.Println("-------------", userName)
	var users []models.User
	db.GetDBConn().Where("first_name = ?", userName).Find(&users)
	context.JSON(http.StatusOK, users)
}
func userStartWith(context *gin.Context) {
	firstLetter := context.Query("firstLetter")


	fmt.Println("-------------", firstLetter)
	var users []models.User
	db.GetDBConn().Where("upper(first_name) LIKE ?", strings.ToUpper(firstLetter) + "%").Find(&users)
	context.JSON(http.StatusOK, users)
}


