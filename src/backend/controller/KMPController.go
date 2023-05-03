package controller

import (
	"backend/algorithm"
	db2 "backend/db"
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func KMPController(c *gin.Context, chat string, idHistory string) {
	var questAns, temp []models.QuestAns
	var chatFromBot models.Chat

	db, err := db2.GetDatabase()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("quest_ans").Find(&questAns); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error.Error(),
		})
		return
	}

	for _, value := range questAns {
		idx, similiarity := algorithm.KMP(chat, value.Question)

		if idx != -1 {
			chatFromBot = models.Chat{
				IdChat:        uuid.New().String(),
				IdHistoryChat: idHistory,
				From:          "bot",
				Chat:          value.Answer,
				Type:          "KMP",
				Time:          time.Now().Local().String(),
			}

			if err := db.Create(chatFromBot); err.Error != nil {
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

		if similiarity > 90 {
			chatFromBot = models.Chat{
				IdChat:        uuid.New().String(),
				IdHistoryChat: idHistory,
				From:          "bot",
				Chat:          value.Answer,
				Type:          "KMP",
				Time:          time.Now().Local().String(),
			}

			if err := db.Create(chatFromBot); err.Error != nil {
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

		if similiarity < 80 {
			temp = append(temp, value)
		}
	}

	if len(temp) != 0 {
		var similar []string

		for idx, value := range temp {
			similar = append(similar, string(idx+1)+". "+value.Question)
		}

		result := strings.Join(similar, "\n")

		chatFromBot = models.Chat{
			IdChat:        uuid.New().String(),
			IdHistoryChat: idHistory,
			From:          "bot",
			Chat:          result,
			Type:          "KMP",
			Time:          time.Now().Local().String(),
		}

		if err := db.Create(chatFromBot); err.Error != nil {
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

	chatFromBot = models.Chat{
		IdChat:        uuid.New().String(),
		IdHistoryChat: idHistory,
		From:          "bot",
		Chat:          "pertanyaan anda tidak ada di database",
		Type:          "KMP",
		Time:          time.Now().Local().String(),
	}

	if err := db.Create(chatFromBot); err.Error != nil {
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
