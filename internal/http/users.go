package http

import (
	"net/http"

	"github.com/JerryJeager/Symptomify-Backend/internal/service/users"
	"github.com/gin-gonic/gin"
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
	if err := ctx.ShouldBindJSON(&verifyUserReq); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required fields",
		})
		return 
	}

	if err := c.serv.VerifyUser(ctx, &verifyUserReq); err != nil{
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, ""))
		return
	}

	ctx.Status(http.StatusOK)
}
