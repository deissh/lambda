package manager

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// id - id контейнера
func (c *Core) Stop(id string) error {
	// todo: добавить поиск по лейблам
	return c.client.ContainerRemove(
		context.Background(),
		id,
		types.ContainerRemoveOptions{})
}
