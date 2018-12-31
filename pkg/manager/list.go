package manager

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"log"
)

func (m *Core) GetAll() ([]types.Container, error) {
	res, err := m.client.ContainerList(context.Background(), types.ContainerListOptions{})
	return res, err
}

func (m *Core) GetActive() ([]types.Container, error) {
	containers, err := m.client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	var res []types.Container
	for _, container := range containers{
		if _, ok := container.Labels["LAMBDA_UUID"]; ok {
			res = append(res, container)
		}
	}

	return res, err
}