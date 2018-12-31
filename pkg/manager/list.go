package manager

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func (m *Core) GetAll() ([]types.Container, error) {
	res, err := m.client.ContainerList(context.Background(), types.ContainerListOptions{})
	return res, err
}
