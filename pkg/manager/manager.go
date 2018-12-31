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

func Create() (managerCore, error) {
	m := managerCore{}
	err := m.newClient()

	return m, err
}

func (m managerCore) newClient () error {
	var err error
	m.client, err = client.NewEnvClient()

	if err != nil {
		log.Fatal(err)
	}

	return err
}
