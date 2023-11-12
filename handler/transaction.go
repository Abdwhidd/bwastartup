package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.CampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Succes get campaign transaction", http.StatusOK, "succes", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
