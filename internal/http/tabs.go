package http

import (
	"net/http"

	"github.com/JerryJeager/Symptomify-Backend/internal/service/tabs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TabController struct {
	serv tabs.TabSv
}

func NewTabController(serv tabs.TabSv) *TabController {
	return &TabController{serv: serv}
}

func (c *TabController) CreateTab(ctx *gin.Context) {
	id, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	userID := uuid.MustParse(id)

	tabID, err := c.serv.CreateTab(ctx, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, ""))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"tab_id": tabID,
	})
}

func (c *TabController) GetTabs(ctx *gin.Context) {
	id, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	userID := uuid.MustParse(id)

	tabs, err := c.serv.GetTabs(ctx, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, ""))
		return
	}

	ctx.JSON(http.StatusCreated, tabs)
}

func (c *TabController) DeleteTab(ctx *gin.Context) {
	var tabID TabIDPathParam
	if err := ctx.ShouldBindUri(&tabID); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, "tab id in path param should be a uuid"))
		return
	}

	if err := c.serv.DeleteTab(ctx, uuid.MustParse(tabID.TabID)); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, "failed to delete the resource"))
		return
	}

	ctx.Status(http.StatusNoContent)
}
