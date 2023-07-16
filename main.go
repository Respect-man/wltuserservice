package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	svc := gin.Default()
	svc.GET("/", CheckServiceStatus)
	svc.Run(":8099")
}

func CheckServiceStatus(ctx *gin.Context) {
	fmt.Printf("I am up and running!!!")
	ctx.JSON(http.StatusOK, "I am up and running!!!")
}
