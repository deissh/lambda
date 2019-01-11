package server

import (
	"github.com/deissh/lambda/pkg/manager"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func inspectHandler(m manager.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")

		data, err := m.Inspect(uuid)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"inspect": data,
		})

		c.JSON(http.StatusNotFound, gin.H{})
	}
}
