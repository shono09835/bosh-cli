package disk

import (
	bicloud "github.com/shono09835/bosh-cli/v7/cloud"
	biconfig "github.com/shono09835/bosh-cli/v7/config"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type ManagerFactory interface {
	NewManager(bicloud.Cloud) Manager
}

type managerFactory struct {
	diskRepo biconfig.DiskRepo
	logger   boshlog.Logger
}

func NewManagerFactory(
	diskRepo biconfig.DiskRepo,
	logger boshlog.Logger,
) ManagerFactory {
	return &managerFactory{
		diskRepo: diskRepo,
		logger:   logger,
	}
}

func (f *managerFactory) NewManager(cloud bicloud.Cloud) Manager {
	return NewManager(cloud, f.diskRepo, f.logger)
}
