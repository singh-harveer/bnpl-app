package handlers

import (
	"bnpl/service"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddMerchant(c *gin.Context) {
	var merchant service.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var err = h.merchantManager.AddMerchant(context.Background(), &merchant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *Handler) GetMerchantByName(c *gin.Context) {
	var param = c.Param(name)
	if param == "" {
		c.JSON(http.StatusBadRequest, errors.New("missing name"))

		return
	}

	var merchant, err = h.merchantManager.GetMerchantByName(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, merchant)
}

func (h *Handler) GetMerchantDiscount(c *gin.Context) {
	var param = c.Param(name)
	if param == "" {
		c.JSON(http.StatusBadRequest, errors.New("missing name"))

		return
	}

	var discountDetails, err = h.merchantManager.Discount(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, discountDetails)
}
