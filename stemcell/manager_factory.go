package stemcell

import (
	bicloud "github.com/shono09835/bosh-cli/v7/cloud"
	biconfig "github.com/shono09835/bosh-cli/v7/config"
)

type ManagerFactory interface {
	NewManager(bicloud.Cloud) Manager
}

type managerFactory struct {
	repo biconfig.StemcellRepo
}

func NewManagerFactory(repo biconfig.StemcellRepo) ManagerFactory {
	return &managerFactory{
		repo: repo,
	}
}

func (f *managerFactory) NewManager(cloud bicloud.Cloud) Manager {
	return NewManager(f.repo, cloud)
}
