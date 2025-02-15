package instance

import (
	"fmt"
	"time"

	biblobstore "github.com/shono09835/bosh-cli/v7/blobstore"
	bicloud "github.com/shono09835/bosh-cli/v7/cloud"
	bidisk "github.com/shono09835/bosh-cli/v7/deployment/disk"
	bideplmanifest "github.com/shono09835/bosh-cli/v7/deployment/manifest"
	bisshtunnel "github.com/shono09835/bosh-cli/v7/deployment/sshtunnel"
	bivm "github.com/shono09835/bosh-cli/v7/deployment/vm"
	bistemcell "github.com/shono09835/bosh-cli/v7/stemcell"
	biui "github.com/shono09835/bosh-cli/v7/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Manager interface {
	FindCurrent() ([]Instance, error)
	Create(
		jobName string,
		id int,
		deploymentManifest bideplmanifest.Manifest,
		cloudStemcell bistemcell.CloudStemcell,
		eventLoggerStage biui.Stage,
	) (Instance, []bidisk.Disk, error)
	DeleteAll(
		pingTimeout time.Duration,
		pingDelay time.Duration,
		skipDrain bool,
		eventLoggerStage biui.Stage,
	) error
}

type manager struct {
	cloud            bicloud.Cloud
	vmManager        bivm.Manager
	blobstore        biblobstore.Blobstore
	sshTunnelFactory bisshtunnel.Factory
	instanceFactory  Factory
	logger           boshlog.Logger
	logTag           string
}

func NewManager(
	cloud bicloud.Cloud,
	vmManager bivm.Manager,
	blobstore biblobstore.Blobstore,
	sshTunnelFactory bisshtunnel.Factory,
	instanceFactory Factory,
	logger boshlog.Logger,
) Manager {
	return &manager{
		cloud:            cloud,
		vmManager:        vmManager,
		blobstore:        blobstore,
		sshTunnelFactory: sshTunnelFactory,
		instanceFactory:  instanceFactory,
		logger:           logger,
		logTag:           "vmDeployer",
	}
}

func (m *manager) FindCurrent() ([]Instance, error) {
	instances := []Instance{}

	// Only one current instance will exist (for now)
	vm, found, err := m.vmManager.FindCurrent()
	if err != nil {
		return instances, bosherr.WrapError(err, "Finding currently deployed instances")
	}

	if found {
		// TODO: store the name of the job for each instance in the repo, so that we can print it when deleting
		jobName := "unknown"
		instanceID := 0

		instance := m.instanceFactory.NewInstance(
			jobName,
			instanceID,
			vm,
			m.vmManager,
			m.sshTunnelFactory,
			m.blobstore,
			m.logger,
		)
		instances = append(instances, instance)
	}

	return instances, nil
}

func (m *manager) Create(
	jobName string,
	id int,
	deploymentManifest bideplmanifest.Manifest,
	cloudStemcell bistemcell.CloudStemcell,
	eventLoggerStage biui.Stage,
) (Instance, []bidisk.Disk, error) {
	var vm bivm.VM
	stepName := fmt.Sprintf("Creating VM for instance '%s/%d' from stemcell '%s'", jobName, id, cloudStemcell.CID())
	err := eventLoggerStage.Perform(stepName, func() error {
		var err error
		vm, err = m.vmManager.Create(cloudStemcell, deploymentManifest)
		if err != nil {
			return bosherr.WrapError(err, "Creating VM")
		}

		if err = cloudStemcell.PromoteAsCurrent(); err != nil {
			return bosherr.WrapErrorf(err, "Promoting stemcell as current '%s'", cloudStemcell.CID())
		}

		return nil
	})
	if err != nil {
		return nil, []bidisk.Disk{}, err
	}

	instance := m.instanceFactory.NewInstance(jobName, id, vm, m.vmManager, m.sshTunnelFactory, m.blobstore, m.logger)

	if err := instance.WaitUntilReady(eventLoggerStage); err != nil {
		return instance, []bidisk.Disk{}, bosherr.WrapError(err, "Waiting until instance is ready")
	}

	disks, err := instance.UpdateDisks(deploymentManifest, eventLoggerStage)
	if err != nil {
		return instance, disks, bosherr.WrapError(err, "Updating instance disks")
	}

	return instance, disks, err
}

func (m *manager) DeleteAll(
	pingTimeout time.Duration,
	pingDelay time.Duration,
	skipDrain bool,
	eventLoggerStage biui.Stage,
) error {
	instances, err := m.FindCurrent()
	if err != nil {
		return err
	}

	for _, instance := range instances {
		if err = instance.Delete(pingTimeout, pingDelay, skipDrain, eventLoggerStage); err != nil {
			return bosherr.WrapErrorf(err, "Deleting existing instance '%s/%d'", instance.JobName(), instance.ID())
		}
	}
	return nil
}
