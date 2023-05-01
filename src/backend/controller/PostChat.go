package controller

import (
	db2 "backend/db"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostChat(c *gin.Context) {
	var chat models.Chat

	db, err := db2.GetDatabase()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Create(&chat); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

	return
}
