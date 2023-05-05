package controller

import (
	db2 "backend/db"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AddChats(c *gin.Context) {
	var chatHistory models.ChatHistory

	db, err := db2.GetDatabase()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	chatHistory.Id = uuid.New().String()
	chatHistory.UpdateAt = utils.GetJktTimeZone()
	chatHistory.LastChat = ""

	if err := db.Create(chatHistory); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func DeleteChats(c *gin.Context) {
	var chatHistory models.ChatHistory

	db, err := db2.GetDatabase()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("chat_histories").Where("id = ?", c.Param("id")).Delete(chatHistory); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
