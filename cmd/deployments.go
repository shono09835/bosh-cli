package cmd

import (
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	biui "github.com/shono09835/bosh-cli/v7/ui"
)

type DeploymentsCmd struct {
	ui       biui.UI
	director boshdir.Director
}

func NewDeploymentsCmd(ui biui.UI, director boshdir.Director) DeploymentsCmd {
	return DeploymentsCmd{ui: ui, director: director}
}

func (c DeploymentsCmd) Run() error {
	deployments, err := c.director.ListDeployments()
	if err != nil {
		return err
	}

	return DeploymentsTable{deployments, c.ui}.Print()
}
