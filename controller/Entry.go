package controller

import (
	"example/jeco-blog/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserWithUsernameAndPassword(ctx *gin.Context) {
	var newUser models.NewUserRegisterWithUsernameAndPassword

	if error := ctx.BindJSON(&newUser); error != nil {
		log.Fatal(error)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "The Data Request Provided is not correct",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Hello World",
		"data":    newUser,
	})
}
