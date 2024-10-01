package controllers

import "github.com/gin-gonic/gin"

func Pinggame(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong by harish and manish",
	})
}
