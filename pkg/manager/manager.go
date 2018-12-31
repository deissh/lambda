package manager

import (
	"github.com/docker/docker/client"
	"log"
)

type Core struct {
	client *client.Client
}

func Create() (Core, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Panic(err)
	}

	return Core{
		client: cli,
	}, nil
}