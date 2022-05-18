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

func (h *userHandler) IsEmailAvailable(c *gin.Context) {
	input := user.CheckEmailInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		c.JSON(
			http.StatusBadRequest,
			helper.APIResponse("Email Check Failed", http.StatusUnprocessableEntity, "error", errorMessage))

		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.APIResponse("Email Check Failed", http.StatusUnprocessableEntity, "error", err.Error()))
		return
	}

	data := gin.H{"is_available": isEmailAvailable}
	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	c.JSON(http.StatusOK, helper.APIResponse(metaMessage, http.StatusOK, "success", data))

}
