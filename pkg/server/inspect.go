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

		containers, err := m.GetActive()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			log.Println(err)
			return
		}

		// todo: убрать вложеность
		for _, cont := range containers {
			for _, name := range cont.Names {
				// fix docker name
				if name == "/"+uuid {
					data, err := m.Inspect(cont.ID)
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
					return
				}
			}
		}

		c.JSON(http.StatusNotFound, gin.H{})
	}
}
