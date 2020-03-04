package trackerruntime

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gofrs/uuid"
)

func createDBContainer(cli *client.Client, config *ContainerDBConfig) (string, error) {
	ctx := context.Background()
	resp, err := cli.ImagePull(ctx, config.Image, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	defer resp.Close()

	portSet := makeExposedPorts(config.ContainerPort)
	hostConfig := makeHostConfig(config.HostPort, config.ContainerPort)

	containerConfig := &container.Config{
		Image:        config.Image,
		Env:          config.Env,
		Cmd:          config.Command,
		ExposedPorts: portSet,
	}

	containerBody, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, makeContainerDBName())
	if err != nil {
		return "", err
	}
	if err := cli.ContainerStart(ctx, containerBody.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	return containerBody.ID, nil
}

func removeDBContainer(cli *client.Client, id string) error {
	ctx := context.Background()
	return cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
		Force: true,
	})
}

func makeContainerDBName() string {
	uid, _ := uuid.NewV4()
	name := "arc-tracker-db" + uid.String()
	return name
}

func makeExposedPorts(port string) nat.PortSet {
	set := nat.PortSet{}
	set[nat.Port(port)] = struct{}{}
	return set
}

func makeHostConfig(hostPort, containerPort string) *container.HostConfig {
	portMap := nat.PortMap{}
	portMap[nat.Port(containerPort)] = []nat.PortBinding{{HostPort: hostPort}}
	return &container.HostConfig{
		PortBindings: portMap,
	}
}
