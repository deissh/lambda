package manager

import (
	"github.com/deissh/lambda/pkg/typings"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

func (c *Core) Create(info typings.FunctionInfo) (string, error) {
	// получаем контейнер
	//_, err := c.client.ImagePull(
	//	context.Background(),
	//	info.Repository.ImageName,
	//	types.ImagePullOptions{})
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}

	// создаем контейнер и настраиваем сеть
	//todo: добавить сеть для контейнеров
	resp, err := c.client.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: info.Repository.ImageName,
			ExposedPorts: map[nat.Port]struct{}{
				"8080": {},
			},
			Labels: map[string]string{
				"lambda.active": "true",
				"lambda.port":   "8080",
			},
			Env: []string{
				"fproccess=" + info.Runtime.Executor + " " + info.Runtime.Cmd,
			},
		},
		&container.HostConfig{
			AutoRemove: true,
			//PortBindings: map[nat.Port][]nat.PortBinding{
			//	"8080": {
			//		{
			//			HostIP:   "",
			//			HostPort: info.Service.Port,
			//		},
			//	},
			//},
		},
		&network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{
				"overlay": {},
			},
		},
		info.Uuid)
	if err != nil {
		return "", err
	}

	// запускаем контейнер
	err = c.client.ContainerStart(
		context.Background(),
		resp.ID,
		types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	return resp.ID, nil
}
