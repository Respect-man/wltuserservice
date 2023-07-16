package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Respect-man/wltuserservice/model"
)

func CreateUser(ctx *gin.Context) {
	usermode := &model.UserObj{}
	if err := ctx.ShouldBindJSON(usermode); err != nil {
		//check if user exist
		if usermode.Username == "" {
			ctx.JSON(http.StatusBadRequest, "username is required")
		}
		if usermode.Phonenumber == "" {
			ctx.JSON(http.StatusBadRequest, "Phone number is required")
		}

	}

	ctx.JSON(http.StatusOK, "User created successfully!!")
}
