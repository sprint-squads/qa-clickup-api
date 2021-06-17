package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Welcome(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome"})
}
