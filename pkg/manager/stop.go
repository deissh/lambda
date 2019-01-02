package manager

import (
	"golang.org/x/net/context"
)

// id - id контейнера
func (c *Core) Stop(id string) error {
	return c.client.ContainerStop(context.Background(), id, nil)
}
