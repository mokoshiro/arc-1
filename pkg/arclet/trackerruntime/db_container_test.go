package trackerruntime

import (
	"context"
	"testing"

	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func TestCreateDBContainer(t *testing.T) {
	client.NewEnvClient()
	cli, err := client.NewEnvClient()
	if err != nil {
		t.Fatal(err)
	}
	config := &ContainerDBConfig{
		Image: "docker.io/library/mysql:5.7",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=root",
			"MYSQL_DATABASE=test_database",
			"MYSQL_USER=arc",
			"MYSQL_PASSWORD=arc",
			"TZ=Asia/Tokyo",
		},
		Command:       []string{"--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"},
		HostPort:      "11111",
		ContainerPort: "3306",
	}
	id, err := createDBContainer(cli, config)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	timeout := time.Second * 1
	if err := cli.ContainerStop(ctx, id, &timeout); err != nil {
		t.Fatal(err)
	}
	time.Sleep(timeout)
	if err := cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
		Force: false,
	}); err != nil {
		t.Fatal(err)
	}
}
