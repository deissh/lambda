package server

import (
	"github.com/deissh/lambda/pkg/typings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createHandler(c *gin.Context) {
	var body typings.FunctionInfo

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}