package handlers

import (
	"bnpl/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	name = "name"
	uri  = "postgres://bnpluser:pass123@0.0.0.0:5434/bnpldb?sslmode=disable"
)

type Handler struct {
	userManager        service.UserManager
	merchantManager    service.MerchantManager
	transactionManager service.TransactionManager
	reporter           service.Reporter
}

func Registration(router *gin.Engine, handler *Handler) {
	// Add new user.
	router.POST("/users", handler.AddUser)
	// Retrieves user details by user name.
	router.GET("/users/:name", handler.GetUserByName)
	router.PUT("/users/:name/payback", handler.DuePayment)
	// Retrieves use credit limit details
	router.GET("/users/:name/creditlimit", handler.GetUserCreditLimit)
	// Add new merchant.
	router.POST("/merchants", handler.AddMerchant)
	// Retrieves merchant details by user name.
	router.GET("/merchants/:name", handler.GetMerchantByName)
	router.GET("/merchants/:name/discount", handler.GetMerchantByName)
	// New transaction.
	router.POST("/transactions", handler.Transaction)
	// Reports all due amount.
	router.GET("/reports/dues", handler.TotalDue)
	// retreive all users which reach at credit limit.
	router.GET("/reports/creditlimits", handler.AllUserAtCreditLimit)
}

// NewHandler created new handler with all dependecy injected.
func NewHandler(ctx context.Context,
	userManagerFunc func(context.Context, string) (service.UserManager, error),
	merchantManagerFunc func(context.Context, string) (service.MerchantManager, error),
	transactiontManagerFunc func(context.Context, string) (service.TransactionManager, error),
	reporterFunc func(context.Context, string) (service.Reporter, error),
) (Handler, error) {
	var userMngr, err = userManagerFunc(ctx, uri)
	if err != nil {
		return Handler{}, fmt.Errorf("failed to create user manager object: %w", err)
	}

	var merchantMngr service.MerchantManager
	merchantMngr, err = merchantManagerFunc(ctx, uri)
	if err != nil {
		return Handler{}, fmt.Errorf("failed to create merchant manager object: %w", err)
	}

	var transactionMngr service.TransactionManager
	transactionMngr, err = transactiontManagerFunc(ctx, uri)
	if err != nil {
		return Handler{}, fmt.Errorf("failed to create transaction manager object: %w", err)
	}

	var reportMngr service.Reporter
	reportMngr, err = reporterFunc(ctx, uri)
	if err != nil {
		return Handler{}, fmt.Errorf("failed to create report manager object: %w", err)
	}

	return Handler{
		userManager:        userMngr,
		merchantManager:    merchantMngr,
		transactionManager: transactionMngr,
		reporter:           reportMngr,
	}, nil
}
