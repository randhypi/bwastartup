package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", err.Error()))
		return
	}

	userFormatter := user.NewUserFormatter(newUser, token)

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

	token, err := h.authService.GenerateToken(loggedInUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", err.Error()))
		return
	}

	userFormatter := user.NewUserFormatter(loggedInUser, token)

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

func (h *userHandler) UploadAvatar(c *gin.Context) {

	// validate token and get user_id
	currentUser := c.MustGet("currentUser").(user.User)
	ID := currentUser.ID

	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false, "error": err.Error()}
		c.JSON(http.StatusBadRequest, helper.APIResponse("Upload Avatar Failed", http.StatusBadRequest, "error", data))
		return
	}

	path := fmt.Sprintf("images/%d-%s", ID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false, "error": err.Error()}
		c.JSON(http.StatusBadRequest, helper.APIResponse("Upload Avatar Failed", http.StatusBadRequest, "error", data))

		return
	}

	_, err = h.userService.SaveAvatar(ID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		c.JSON(http.StatusUnprocessableEntity, helper.APIResponse("Upload Avatar Failed", http.StatusUnprocessableEntity, "error", data))

		return
	}

	data := gin.H{"is_uploaded": true}

	c.JSON(http.StatusOK, helper.APIResponse("Upload Avatar Success", http.StatusOK, "success", data))
}
