package docker

import (
	"go.k6.io/k6/js/modules"
)

const version = "v0.0.1"

func init() {
	modules.Register("k6/x/docker", &Docker{
		Version: version,
	})

	containers := Containers{Version: version}
	containers.SetupClient()
	modules.Register("k6/x/docker/containers", &containers)

	volumes := Volumes{Version: version}
	volumes.SetupClient()
	modules.Register("k6/x/docker/volumes", &volumes)

	networks := Networks{Version: version}
	networks.SetupClient()
	modules.Register("k6/x/docker/networks", &networks)

	images := Images{Version: version}
	images.SetupClient()
	modules.Register("k6/x/docker/images", &images)
}

// Docker is the main export of k6 docker extension
type Docker struct {
	Version string
}
