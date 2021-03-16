package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type Volumes struct {
	Version string
	Client  *client.Client
}

// SetupClient for filling in Docker client instance
func (d *Volumes) SetupClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	d.Client = cli
}

// Create works as Docker volume create
func (v *Volumes) Create(options volumetypes.VolumeCreateBody) (types.Volume, error) {
	return v.Client.VolumeCreate(context.Background(), options)
}

// List works as Docker volume ls
func (v *Volumes) List(filter filters.Args) (volumetypes.VolumeListOKBody, error) {
	return v.Client.VolumeList(context.Background(), filter)
}

// Remove works as Docker volume rm
func (v *Volumes) Remove(volumeID string, force bool) error {
	return v.Client.VolumeRemove(context.Background(), volumeID, force)
}
