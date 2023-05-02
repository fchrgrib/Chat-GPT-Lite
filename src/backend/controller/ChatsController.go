package controller

import (
	"backend/algorithm"
	db2 "backend/db"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func PostChatsController(c *gin.Context) {
	var chatFromUser models.Chat
	var chatFromBot models.Chat

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
	chatFromUser.Time = time.Now().Local().String()

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

	if isCalculate, numbers := utils.CheckCalculate(chatFromUser.Chat); isCalculate {
		chatFromBot = models.Chat{
			IdChat:        uuid.New().String(),
			IdHistoryChat: chatFromUser.IdHistoryChat,
			From:          "bot",
			Chat:          "Hasilnya adalah " + (algorithm.Calculate(numbers)),
			Type:          chatFromUser.Type,
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

	if isDate, date := utils.CheckDate(chatFromUser.Chat); isDate {
		chatFromBot = models.Chat{
			IdChat:        uuid.New().String(),
			IdHistoryChat: chatFromUser.IdHistoryChat,
			From:          "bot",
			Chat:          "Tanggal tersebut adalah hari " + (utils.SearchDay(date)),
			Type:          chatFromUser.Type,
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

	if isInsert, isSame, matches1, matches2 := utils.CheckInsertQuesAns(chatFromUser.Chat); isInsert {
		var questAns models.QuestAns
		if isSame {

			_ = db.Where("question = ?", matches1).First(&questAns)
			questAns.Answer = matches2
			questAns.TimeCreate = time.Now().Local().String()
			db.Save(questAns)

			chatFromBot = models.Chat{
				IdChat:        uuid.New().String(),
				IdHistoryChat: chatFromUser.IdHistoryChat,
				From:          "bot",
				Chat:          matches1 + " sudah ada pada database, tetapi jawaban akan diganti dengan " + matches2,
				Type:          chatFromUser.Type,
				Time:          time.Now().Local().String(),
			}

			if err := db.Create(chatFromBot); err.Error != nil {
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
			Time:          time.Now().Local().String(),
		}

		if err := db.Create(chatFromBot); err.Error != nil {
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
			IdHistoryChat: chatFromUser.IdHistoryChat,
			From:          "bot",
			Chat:          "Pertanyaan " + question + " telah dihapus",
			Type:          chatFromUser.Type,
			Time:          time.Now().Local().String(),
		}

		if err := db.Create(chatFromBot); err.Error != nil {
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

	//TODO make reply chat using KM and BM

	return
}

//TODO make GET request to get all chats
