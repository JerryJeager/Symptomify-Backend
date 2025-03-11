package http

import (
	"net/http"

	"github.com/JerryJeager/Symptomify-Backend/internal/service/chats"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatController struct {
	serv chats.ChatSv
}

func NewChatController(serv chats.ChatSv) *ChatController {
	return &ChatController{serv: serv}
}

func (c *ChatController) CreateChat(ctx *gin.Context) {
	id, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	userID := uuid.MustParse(id)

	var chat chats.Chat
	if err := ctx.ShouldBindJSON(&chat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required fields",
		})
		return
	}

	var tabID TabIDPathParam
	if err := ctx.ShouldBindUri(&tabID); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, "tab id in path param should be a uuid"))
		return
	}

	chatId, err := c.serv.CreateChat(ctx, userID, uuid.MustParse(tabID.TabID), &chat)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"chat_id": chatId,
	})
}

func (c *ChatController) GetChatByTabID(ctx *gin.Context) {
	var tabID TabIDPathParam
	if err := ctx.ShouldBindUri(&tabID); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, "tab id in path param should be a uuid"))
		return
	}

	chats, err := c.serv.GetChatByTabID(ctx, uuid.MustParse(tabID.TabID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, chats)
}

func (c *ChatController) DeleteChat(ctx *gin.Context) {
	var chatID ChatIDPathParam
	if err := ctx.ShouldBindUri(&chatID); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, "chat id in path param should be a uuid"))
		return
	}

	err := c.serv.DeleteChat(ctx, uuid.MustParse(chatID.ChatID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	ctx.Status(http.StatusNoContent)
}
