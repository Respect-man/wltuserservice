package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	usermode := &model.UserObj{}
	readRequest := ctx.ShouldBindJSON(usermode)
	if readRequest != null {

	}

	ctx.JSON(http.StatusOK, "User created successfully!!")
}
