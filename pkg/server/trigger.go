package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func triggerHandler(c *gin.Context) {
	uuid := c.Param("uuid")[0:12]

	// todo: add read labels
	proxyurl, err := url.Parse("http://" + uuid + ":8080")
	log.Println("proxy to " + proxyurl.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	proxy := httputil.NewSingleHostReverseProxy(proxyurl)
	proxy.ServeHTTP(c.Writer, c.Request)
}
