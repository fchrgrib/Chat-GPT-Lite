package controller

import (
	"backend/algorithm"
	db2 "backend/db"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func PostChatsController(c *gin.Context) {
	var chatFromUser models.Chat
	var chatFromBot models.Chat
	var chatHistory models.ChatHistory

	db, err := db2.GetDatabase()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	chatFromUser.IdHistoryChat = c.Param("id")
	chatFromUser.From = "user"
	chatFromUser.IdChat = uuid.New().String()
	chatFromUser.Time = utils.GetJktTimeZone()

	if err := c.ShouldBindJSON(&chatFromUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Create(chatFromUser); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error.Error(),
		})
		return
	}

	if isDate, date := utils.CheckDate(chatFromUser.Chat); isDate {
		chatFromBot = models.Chat{
			IdChat:        uuid.New().String(),
			IdHistoryChat: chatFromUser.IdHistoryChat,
			From:          "bot",
			Chat:          "Tanggal tersebut adalah hari " + (utils.SearchDay(date)),
			Type:          chatFromUser.Type,
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

	if isCalculate, numbers := utils.CheckCalculate(chatFromUser.Chat); isCalculate {
		chatFromBot = models.Chat{
			IdChat:        uuid.New().String(),
			IdHistoryChat: chatFromUser.IdHistoryChat,
			From:          "bot",
			Chat:          "Hasilnya adalah " + (algorithm.Calculate(numbers)),
			Type:          chatFromUser.Type,
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

	if isInsert, isSame, matches1, matches2 := utils.CheckInsertQuesAns(chatFromUser.Chat); isInsert {
		var questAns models.QuestAns
		if isSame {

			_ = db.Where("question = ?", matches1).First(&questAns)
			questAns.Answer = matches2
			questAns.TimeCreate = utils.GetJktTimeZone()
			db.Save(questAns)

			chatFromBot = models.Chat{
				IdChat:        uuid.New().String(),
				IdHistoryChat: chatFromUser.IdHistoryChat,
				From:          "bot",
				Chat:          matches1 + " sudah ada pada database, tetapi jawaban akan diganti dengan " + matches2,
				Type:          chatFromUser.Type,
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

			if err := utils.UpdateAnswer(matches1, matches2); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": err.Error(),
				})
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
			Chat:          "Pertanyaan " + matches1 + " telah ditambahkan",
			Type:          chatFromUser.Type,
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

		if err := utils.InsertQuestionAnswer(matches1, matches2); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		return
	}

	if isAsk, isFound, question := utils.CheckEraseQues(chatFromUser.Chat); isAsk {

		if !isFound {
			chatFromBot = models.Chat{
				IdChat:        uuid.New().String(),
				IdHistoryChat: chatFromUser.IdHistoryChat,
				From:          "bot",
				Chat:          "Pertanyaan " + question + " tidak ditemukan di database",
				Type:          chatFromUser.Type,
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

		chatFromBot = models.Chat{
			IdChat:        uuid.New().String(),
			IdHistoryChat: chatFromUser.IdHistoryChat,
			From:          "bot",
			Chat:          "Pertanyaan " + question + " telah dihapus",
			Type:          chatFromUser.Type,
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

		var questAns models.QuestAns

		if err := db.Where("question = ?", question).Delete(questAns); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		return
	}

	if chatFromUser.Type == "KMP" {
		KMPController(c, chatFromUser)
		return
	}
	if chatFromUser.Type == "BM" {
		BMController(c, chatFromUser)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "BadRequest",
	})

	return
}

func GetChatsController(c *gin.Context) {
	chats, _ := utils.GetChats(c.Param("id"))
	history, _ := utils.GetHistoryChats()

	c.JSON(http.StatusOK, gin.H{
		"id_history": c.Param("id"),
		"chats":      chats,
		"history":    history,
	})

	return
}

func FirstPage(c *gin.Context) {
	var chats []models.Chat
	history, _ := utils.GetHistoryChats()

	c.JSON(http.StatusOK, gin.H{
		"id_history": c.Param("id"),
		"chats":      chats,
		"history":    history,
	})

	return
}
