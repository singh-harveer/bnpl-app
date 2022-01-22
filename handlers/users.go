package handlers

import (
	"bnpl/service"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddUser(c *gin.Context) {
	var user service.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	log.Println(user)
	var err = h.userManager.AddUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *Handler) GetUserByName(c *gin.Context) {
	var param = c.Param(name)
	if param == "" {
		c.JSON(http.StatusBadRequest, errors.New("missing name"))

		return
	}

	var user, err = h.userManager.GetUserByName(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUserCreditLimit(c *gin.Context) {
	var param = c.Param(name)
	if param == "" {
		c.JSON(http.StatusBadRequest, errors.New("missing name"))

		return
	}

	var creditDetails, err = h.userManager.CreditLimit(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, creditDetails)
}
