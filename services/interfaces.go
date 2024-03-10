package services

import (
	"github.com/chat-bot-data/models"
	"github.com/gin-gonic/gin"
)

type User interface {
	Create(c *gin.Context, data models.UserInfo) (models.UserInfo, error)
	GetByID(c *gin.Context, ID string) (models.UserInfo, error)
	Get(c *gin.Context, username string, password string) (models.UserInfo, error)
	PatchByID(c *gin.Context, ID string, user models.UserInfo) (models.UserInfo, error)
}
