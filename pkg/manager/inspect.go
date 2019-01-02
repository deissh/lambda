package manager

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func (m *Core) Inspect(id string) (types.ContainerJSON, error) {
	res, err := m.client.ContainerInspect(context.Background(), id)
	return res, err
}
