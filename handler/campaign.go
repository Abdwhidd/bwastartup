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
		var errors = helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed to create campaign", http.StatusBadRequest, "error", errorMessage)
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
		response := helper.ApiResponse("Failed to create campaign bydata", http.StatusBadRequest, "error", nil)
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

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		var errors = helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed to upload campaign image:", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !input.IsPrimary {
		errorMessage := gin.H{"errors": "IsPrimary field is required and must be true"}
		response := helper.ApiResponse("Failed to upload campaign image:", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		handleUploadError(c, "Failed to upload campaign image")
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		handleUploadError(c, "Failed to upload campaign image")
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		handleUploadError(c, "Failed to upload campaign image")
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Success upload campaign images", http.StatusOK, "errors", data)
	c.JSON(http.StatusOK, response)
}

func handleUploadError(c *gin.Context, errorMsg string) {
	data := gin.H{"is_uploaded": false}
	response := helper.ApiResponse(errorMsg, http.StatusBadRequest, "errors", data)
	c.JSON(http.StatusBadRequest, response)
}
