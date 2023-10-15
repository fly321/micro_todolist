package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var userRequest service.UserRequest
	PanicIfUserError(ctx.Bind(&userRequest))
	// 从gin.keys中获取userService
	userService := ctx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userRequest)
	PanicIfUserError(err)
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": userResp,
	})
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var userRequest service.UserRequest
	PanicIfUserError(ctx.Bind(&userRequest))
	// 从gin.keys中获取userService
	userService := ctx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userRequest)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	ctx.JSON(200, gin.H{
		"code": userResp.Code,
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
		"msg": "登录成功",
	})
}
