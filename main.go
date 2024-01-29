package main

import (
	"fmt"
	"net/http"

	ctrn "github.com/Respect-man/wltuserservice/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DbGorm *gorm.DB
var ErrGorm error

func main() {

	svc := gin.Default()
	svc.GET("/", CheckServiceStatus)
	svc.POST("/user/createuser", ctrn.CreateUser)
	svc.Run(":8080")
}

func CheckServiceStatus(ctx *gin.Context) {
	fmt.Printf("I am up and running!!!")
	ctx.JSON(http.StatusOK, "I am up and running!!!")
}
