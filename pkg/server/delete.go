package server

import (
	"github.com/deissh/lambda/pkg/manager"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func deleteHandler(m manager.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")

		if err := m.Stop(uuid); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "function uuid not founded",
			})
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}
