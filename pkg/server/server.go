package server

import (
	"github.com/deissh/lambda/pkg/manager"
	"github.com/gin-gonic/gin"
)

type core struct {
	ip      string
	port    string
	router  *gin.Engine
	manager manager.Core
}

func Run(host string, port string) error {
	s := core{
		ip:   host,
		port: port,
	}
	return s.Start()
}

func (s core) Start() error {
	s.manager = manager.Create()
	s.router = gin.Default()

	// create routing
	s.Routing()
	s.router.Use(gin.Recovery())

	return s.router.Run(s.ip + ":" + s.port)
}

func (s core) Routing() {

	g := s.router.Group("/v1")
	{
		g.GET("/")
		g.POST("/create", createHandler(s.manager))
		g.Any("/function/:uuid", triggerHandler)
		g.GET("/delete/:uuid", deleteHandler(s.manager))
		g.GET("/inspect/:uuid", inspectHandler(s.manager))
	}
}
