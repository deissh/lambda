package manager

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"log"
)

type managerCore struct {
	context context.Context
	client *client.Client
}

func Create() managerCore {
	m := managerCore{}
	return m
}

func (m managerCore) NewClient () error {
	var err error
	m.client, err = client.NewEnvClient()

	if err != nil {
		log.Fatal(err)
	}

	return err
}
