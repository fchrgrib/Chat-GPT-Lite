package controller

import (
	"backend/algorithm"
	db2 "backend/db"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
)

func BMController(c *gin.Context, chatFromUser models.Chat) {
	var questAns, temp []models.QuestAns
	var chatFromBot models.Chat
	var chatHistory models.ChatHistory

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

	limitSimilar := 0
	for _, value := range questAns {
		idx, similarity := algorithm.BM(chatFromUser.Chat, value.Question)

		if idx != -1 && similarity == 100 {
			chatFromBot = models.Chat{
				IdChat:        uuid.New().String(),
				IdHistoryChat: chatFromUser.IdHistoryChat,
				From:          "bot",
				Chat:          value.Answer,
				Type:          "BM",
				Time:          utils.GetJktTimeZone(),
			}

			if err := db.Create(chatFromBot); err.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": err.Error.Error(),
				})
				return
			}
			if err := db.Table("chat_histories").Where("id", chatFromUser.IdHistoryChat).First(&chatHistory); err.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": err.Error.Error(),
				})
				return
			}
			chatHistory.LastChat = chatFromBot.Chat

			chatHistory.UpdateAt = utils.GetJktTimeZone()

			if err := db.Save(chatHistory); err.Error != nil {
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

		if similarity >= 90 {
			chatFromBot = models.Chat{
				IdChat:        uuid.New().String(),
				IdHistoryChat: chatFromUser.IdHistoryChat,
				From:          "bot",
				Chat:          value.Answer,
				Type:          "BM",
				Time:          utils.GetJktTimeZone(),
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

		if similarity < 90 && idx != -1 && limitSimilar < 3 {
			temp = append(temp, value)
			limitSimilar++
		}
	}

	if len(temp) != 0 {
		var similar []string
		similar = append(similar, "maksud anda :")

		for idx, value := range temp {
			similar = append(similar, strconv.Itoa(idx+1)+". "+value.Question)
		}

		result := strings.Join(similar, "\n")

		chatFromBot = models.Chat{
			IdChat:        uuid.New().String(),
			IdHistoryChat: chatFromUser.IdHistoryChat,
			From:          "bot",
			Chat:          result,
			Type:          "BM",
			Time:          utils.GetJktTimeZone(),
		}
		if err := db.Table("chat_histories").Where("id = ?", chatFromUser.IdHistoryChat).First(&chatHistory); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error.Error(),
			})
			return
		}
		chatHistory.LastChat = chatFromBot.Chat

		chatHistory.UpdateAt = utils.GetJktTimeZone()

		if err := db.Save(chatHistory); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error.Error(),
			})
			return
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
		IdHistoryChat: chatFromUser.IdHistoryChat,
		From:          "bot",
		Chat:          "pertanyaan anda tidak ada di database",
		Type:          "BM",
		Time:          utils.GetJktTimeZone(),
	}
	if err := db.Table("chat_histories").Where("id = ?", chatFromUser.IdHistoryChat).First(&chatHistory); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error.Error(),
		})
		return
	}
	chatHistory.LastChat = chatFromBot.Chat
	chatHistory.UpdateAt = utils.GetJktTimeZone()

	if err := db.Save(chatHistory); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error.Error(),
		})
		return
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
