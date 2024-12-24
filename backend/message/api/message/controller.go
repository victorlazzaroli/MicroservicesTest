package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/victorlazzaroli/microservicesTest/auth/api/message/dto"
	"gorm.io/gorm"
)

type messageController struct {
	service MessageServiceI
}

type MessageControllerI interface {
	GetMessage(ctx *gin.Context)
	GetAllMessages(ctx *gin.Context)
	GetAllAuthorMessages(ctx *gin.Context)
	SaveMessage(ctx *gin.Context)
	UpdateMessage(ctx *gin.Context)
	DeleteMessage(ctx *gin.Context)
}

func (m *messageController) UpdateMessage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var message dto.UpdateMessageRequest
	if err := ctx.ShouldBindBodyWithJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if errore := m.service.UpdateMessage(id, message); errore != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errore})

		return
	}
}

// DeleteMessage implements MessageControllerI.
func (m *messageController) DeleteMessage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if errore := m.service.DeleteMessage(id); errore != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errore.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// GetAllAuthorMessages implements MessageControllerI.
func (m *messageController) GetAllAuthorMessages(ctx *gin.Context) {
	authorId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	messages, errore := m.service.GetMessages(authorId)
	if errore != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errore.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

// GetAllMessages implements MessageControllerI.
func (m *messageController) GetAllMessages(ctx *gin.Context) {
	messages, errore := m.service.GetMessages(-1)
	if errore != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errore.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

// GetMessage implements MessageControllerI.
func (m *messageController) GetMessage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	message, errore := m.service.GetMessage(id)
	if errore != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errore.Error()})
		return
	}

	ctx.JSON(http.StatusOK, message)
}

// SaveMessage implements MessageControllerI.
func (m *messageController) SaveMessage(ctx *gin.Context) {
	var message dto.CreateMessageRequest
	if err := ctx.ShouldBindBodyWithJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := m.service.CreateMessage(&message); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func NewMessageController(engine *gin.Engine, db *gorm.DB) MessageControllerI {
	repository := NewMessageRepository(db)
	service := NewMessageService(&repository)
	controller := &messageController{
		service: service,
	}

	api := engine.Group("messages")
	{
		api.GET("author/:id", controller.GetAllAuthorMessages)

		api.GET("/", controller.GetAllMessages)
		api.POST("/", controller.SaveMessage)

		api.GET("/:id", controller.GetMessage)
		api.DELETE("/:id", controller.DeleteMessage)
		api.PATCH("/:id", controller.UpdateMessage)
	}

	return controller
}
