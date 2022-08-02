package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)

	if err != nil {
		response := helper.APIResponse(
			"Error to get campaign",
			http.StatusBadRequest, "error",
			campaign.FormatCampaigns(campaigns),
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "Success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	input := campaign.GetCampaignDetailInput{}
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Error to get campaign",
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignById(input)

	if err != nil {
		response := helper.APIResponse(
			"Error to get campaign",
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign", http.StatusOK, "Success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	input := campaign.CreateCampaignInput{}
	input.User = currentUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Error to create campaign",
			http.StatusUnprocessableEntity, "error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCampaign, err := h.service.CreateCampaign(input)

	if err != nil {
		response := helper.APIResponse(
			"Error to create campaign",
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign", http.StatusOK, "Success", campaign.FormatCampaignDetail(newCampaign))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {

	inputID := campaign.GetCampaignDetailInput{}
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse(
			"Error to update campaign",
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	inputData := campaign.CreateCampaignInput{}
	inputData.User = currentUser
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Error to update campaign",
			http.StatusUnprocessableEntity, "error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)

	if err != nil {
		response := helper.APIResponse(
			"Error to update campaign",
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign", http.StatusOK, "Success", campaign.FormatCampaignDetail(updatedCampaign))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	input := campaign.CreateCampaignImageInput{}
	input.User = currentUser
	err := c.ShouldBind(&input)

	fmt.Println(input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Error to upload image",
			http.StatusUnprocessableEntity, "error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("image")

	if err != nil {
		data := gin.H{"is_uploaded": false, "error": err.Error()}
		c.JSON(http.StatusBadRequest, helper.APIResponse("Upload Image Failed", http.StatusBadRequest, "error", data))
		return
	}

	path := fmt.Sprintf("images/campaigns/%d-%d-%s", userID, input.CampaignID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false, "error": err.Error()}
		c.JSON(http.StatusBadRequest, helper.APIResponse("Upload Image Failed", http.StatusBadRequest, "error", data))
		return
	}
	image, err := h.service.SaveCampaignImage(input, path)

	if err != nil {
		response := helper.APIResponse(
			err.Error(),
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true, "image": image}

	response := helper.APIResponse("Image", http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)

}
