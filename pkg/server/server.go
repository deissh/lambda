package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

type core struct {
	ip string
	port string
	router *gin.Engine
}

func Run(host string, port string) error {
	s := core{
		ip:   host,
		port: port,
	}
	return s.Start()
}

func (s core) Start() error {
	s.router = gin.Default()

	log.Println("starting server")

	// create routing
	s.Routing()
	s.router.Use(gin.Recovery())

	return s.router.Run(s.ip + ":" + s.port)
}

func (s core) Routing() {
	g := s.router.Group("/v1")
	{
		g.GET("/")
		g.POST("/create", createHandler)
		g.Any("/function/:uuid")
	}
}