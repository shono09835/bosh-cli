package deployment

import (
	bidisk "github.com/shono09835/bosh-cli/v7/deployment/disk"
	biinstance "github.com/shono09835/bosh-cli/v7/deployment/instance"
	bistemcell "github.com/shono09835/bosh-cli/v7/stemcell"
	biui "github.com/shono09835/bosh-cli/v7/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Manager interface {
	FindCurrent() (deployment Deployment, found bool, err error)
	Cleanup(biui.Stage) error
}

type manager struct {
	instanceManager   biinstance.Manager
	diskManager       bidisk.Manager
	stemcellManager   bistemcell.Manager
	deploymentFactory Factory
}

func NewManager(
	instanceManager biinstance.Manager,
	diskManager bidisk.Manager,
	stemcellManager bistemcell.Manager,
	deploymentFactory Factory,
) Manager {
	return &manager{
		instanceManager:   instanceManager,
		diskManager:       diskManager,
		stemcellManager:   stemcellManager,
		deploymentFactory: deploymentFactory,
	}
}

func (m *manager) FindCurrent() (deployment Deployment, found bool, err error) {
	instances, err := m.instanceManager.FindCurrent()
	if err != nil {
		return nil, false, bosherr.WrapError(err, "Finding current deployment instances")
	}

	disks, err := m.diskManager.FindCurrent()
	if err != nil {
		return nil, false, bosherr.WrapError(err, "Finding current deployment disks")
	}

	stemcells, err := m.stemcellManager.FindCurrent()
	if err != nil {
		return nil, false, bosherr.WrapError(err, "Finding current deployment stemcells")
	}

	if len(instances) == 0 && len(disks) == 0 && len(stemcells) == 0 {
		return nil, false, nil
	}

	return m.deploymentFactory.NewDeployment(instances, disks, stemcells), true, nil
}

func (m *manager) Cleanup(stage biui.Stage) error {
	if err := m.diskManager.DeleteUnused(stage); err != nil {
		return err
	}

	if err := m.stemcellManager.DeleteUnused(stage); err != nil {
		return err
	}

	return nil
}
