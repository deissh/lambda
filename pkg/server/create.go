package server

import (
	"github.com/deissh/lambda/pkg/manager"
	"github.com/deissh/lambda/pkg/typings"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func createHandler(m manager.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body typings.FunctionInfo

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			log.Println(err.Error())
			return
		}

		uuid, err := m.Create(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"uuid": uuid,
		})
	}
}
