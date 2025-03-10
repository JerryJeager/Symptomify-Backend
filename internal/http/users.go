package http

import (
	"net/http"

	"github.com/JerryJeager/Symptomify-Backend/internal/service/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	serv users.UserSv
}

func NewUserController(serv users.UserSv) *UserController {
	return &UserController{serv: serv}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required fields",
		})
		return
	}

	if err := c.serv.CreateUser(ctx, &user); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, ""))
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *UserController) VerifyUser(ctx *gin.Context) {
	var verifyUserReq users.VerifyUserReq
	if err := ctx.ShouldBindJSON(&verifyUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required fields",
		})
		return
	}

	if err := c.serv.VerifyUser(ctx, &verifyUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, ""))
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginReq users.LoginReq
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required fields",
		})
		return
	}

	token, err := c.serv.Login(ctx, &loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	userID := uuid.MustParse(id)

	user, err := c.serv.GetUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetErrorJson(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": user.ID,
		"email": user.Email, 
		"name": user.Name,
		"created_at": user.CreatedAt,
		"is_verified": user.IsVerified,
	})
}
