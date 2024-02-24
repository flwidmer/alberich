package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ListContainers lists all containers running on the Docker host.
func ListContainers() error {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	fmt.Println("Containers running:")
	for _, container := range containers {
		fmt.Printf("ID: %s Image: %s Command: %q Created: %s Status: %s\n",
			container.ID[:10], container.Image, container.Command, container.Created, container.Status)
	}

	return nil
}
