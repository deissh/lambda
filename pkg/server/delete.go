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

		containers, err := m.GetActive()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			log.Println(err)
			return
		}

		for _, cont := range containers {
			for _, name := range cont.Names {
				// fix docker name
				if name == "/"+uuid {
					if err := m.Stop(cont.ID); err != nil {
						log.Println(err)
						c.JSON(http.StatusInternalServerError, gin.H{
							"error": err.Error(),
						})
					}

					c.JSON(http.StatusOK, gin.H{})
				}
			}
		}
	}
}
