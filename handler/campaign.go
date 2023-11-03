package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
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

	campaigns, err := h.service.FindCampaign(userId)
	if err != nil {
		response := helper.ApiResponse("Failed find of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List Of Campaign", http.StatusOK, "error", campaign.FormatterCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.CampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed find of campaign 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.FindCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Failed find of campaign 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign Detail", http.StatusOK, "succes", campaign.FormatterCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CampaignCreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	createCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.ApiResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Create campaign", http.StatusOK, "succes", campaign.FormatterCampaign(createCampaign))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UpdatedCampaign(c *gin.Context) {
	var inputId campaign.CampaignDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.ApiResponse("Failed Update campaign byid", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CampaignCreateInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to create campaign bydata", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to create campaign server", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Update campaign", http.StatusOK, "succes", campaign.FormatterCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}
