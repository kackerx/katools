package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kackerx/learngin/models"
	"net/http"
)

type VideoController interface {
	GetAll(context *gin.Context)
}

type controller struct {
	videos []models.Video
}

func (c *controller) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func NewController() *controller {
	return &controller{}
}
