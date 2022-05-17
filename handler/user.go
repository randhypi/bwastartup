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
	input := user.RegisterUserInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		c.JSON(
			http.StatusBadRequest,
			helper.APIResponse("Account Failed Registered", http.StatusUnprocessableEntity, "error", errorMessage))

		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", err.Error()))
		return
	}

	userFormatter := user.NewUserFormatter(newUser, "token")

	c.JSON(http.StatusOK, helper.APIResponse("User registered successfully", http.StatusOK, "success", userFormatter))
}

func (h *userHandler) Login(c *gin.Context) {
	input := user.LoginInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		c.JSON(
			http.StatusBadRequest,
			helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage))

		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", err.Error()))
		return
	}

	userFormatter := user.NewUserFormatter(loggedInUser, "token")

	c.JSON(http.StatusOK, helper.APIResponse("Login Success", http.StatusOK, "success", userFormatter))
}
