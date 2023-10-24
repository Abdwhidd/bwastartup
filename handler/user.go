package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita parsing menjadi parameter service

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		var errors = helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Registered account failed", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if true {
		response := helper.ApiResponse("Registered account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatterUser(newUser, "tokeennnnnnnnnnn")
	response := helper.ApiResponse("Account has be registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
