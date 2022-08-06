package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *handler {
	return &handler{service}
}

func (h *handler) GetTransactionsByCampaignID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	input := transaction.GetTransactionsByCampaignIDInput{}
	input.User = currentUser

	err := c.ShouldBindUri(&input)
	if err != nil {

		response := helper.APIResponse(
			"Error to get Transaction by campaign id",
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	transactions, err := h.service.GetTransactionsByCampaignID(input)

	if err != nil {
		response := helper.APIResponse(
			err.Error(),
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	newCampaignTransactions := transaction.NewCampaignTransactionFormatterList(transactions)

	response := helper.APIResponse("List of transactions", http.StatusOK, "Success", newCampaignTransactions)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	userID := currentUser.ID

	transactions, err := h.service.GetTransactionByUserID(userID)

	if err != nil {
		response := helper.APIResponse(
			"Failed to get user's transaction",
			http.StatusBadRequest, "error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	newUserTransactions := transaction.NewUserTransactionFormatterList(transactions)

	response := helper.APIResponse("List of by user transactions", http.StatusOK, "Success", newUserTransactions)
	c.JSON(http.StatusOK, response)
}
