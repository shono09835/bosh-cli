package deployment

import (
	"time"

	bidisk "github.com/shono09835/bosh-cli/v7/deployment/disk"
	biinstance "github.com/shono09835/bosh-cli/v7/deployment/instance"
	bistemcell "github.com/shono09835/bosh-cli/v7/stemcell"
)

type Factory interface {
	NewDeployment(
		[]biinstance.Instance,
		[]bidisk.Disk,
		[]bistemcell.CloudStemcell,
	) Deployment
}

type factory struct {
	pingTimeout time.Duration
	pingDelay   time.Duration
}

func NewFactory(
	pingTimeout time.Duration,
	pingDelay time.Duration,
) Factory {
	return &factory{
		pingTimeout: pingTimeout,
		pingDelay:   pingDelay,
	}
}

func (f *factory) NewDeployment(
	instances []biinstance.Instance,
	disks []bidisk.Disk,
	stemcells []bistemcell.CloudStemcell,
) Deployment {
	return NewDeployment(
		instances,
		disks,
		stemcells,
		f.pingTimeout,
		f.pingDelay,
	)
}
