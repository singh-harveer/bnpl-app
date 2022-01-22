package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TotalDue retrieves all due amount for all users.
func (h *Handler) TotalDue(c *gin.Context) {
	var dueDetails, err = h.reporter.TotalDue(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	log.Println(dueDetails)
	c.JSON(http.StatusOK, dueDetails)
}

// AllUserAtCreditLimit retrieves all users which reached credit limits.
func (h *Handler) AllUserAtCreditLimit(c *gin.Context) {
	var dueDetails, err = h.reporter.AllUserAtCreditLimit(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, dueDetails)
}
