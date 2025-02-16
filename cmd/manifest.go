package cmd

import (
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
)

type ManifestCmd struct {
	ui         boshui.UI
	deployment boshdir.Deployment
}

func NewManifestCmd(ui boshui.UI, deployment boshdir.Deployment) ManifestCmd {
	return ManifestCmd{ui: ui, deployment: deployment}
}

func (c ManifestCmd) Run() error {
	manifest, err := c.deployment.Manifest()
	if err != nil {
		return err
	}

	c.ui.PrintBlock([]byte(manifest))

	return nil
}
