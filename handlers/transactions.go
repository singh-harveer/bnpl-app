package handlers

import (
	"bnpl/service"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type request struct {
	UserName     string  `json:"userName"`
	MerchantName string  `json:"merchantName"`
	Amount       float64 `json:"amount"`
}

type duePayRequest struct {
	UserName string  `json:"userName"`
	Amount   float64 `json:"amount"`
}

func (req *duePayRequest) validate() error {
	var userName = strings.TrimSpace(req.UserName)

	if userName == "" {
		return errors.New("userName is missing")
	}

	if req.Amount < 1 {
		return errors.New("invalid amount")
	}

	return nil
}

func (req *request) validate() error {
	var merchantName = strings.TrimSpace(req.MerchantName)
	var userName = strings.TrimSpace(req.UserName)

	if merchantName == "" || userName == "" {
		return errors.New("username or merchantname is missing")
	}

	if req.Amount < 1 {
		return errors.New("invalid amount")
	}

	return nil
}

// Transaction creates new transactions.
func (h *Handler) Transaction(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := req.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	var ctx = context.Background()

	var user, err = h.userManager.GetUserByName(ctx, req.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to retrieve user details"})

		return
	}

	var merchant service.Merchant
	merchant, err = h.merchantManager.GetMerchantByName(ctx, req.MerchantName)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to retrieve merchant details"})

		return
	}

	err = h.transactionManager.Create(ctx,
		service.Transaction{
			UserID:     user.ID,
			MerchantID: merchant.ID,
			Amount:     req.Amount,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, nil)
}

// DuePayment deposit amount towards user due.
func (h *Handler) DuePayment(c *gin.Context) {
	var req duePayRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := req.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	var result, err = h.userManager.DuePayment(
		context.Background(),
		req.UserName,
		req.Amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, result)
}
