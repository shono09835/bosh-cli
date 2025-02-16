package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
)

type OrphanDiskCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewOrphanDiskCmd(ui boshui.UI, director boshdir.Director) OrphanDiskCmd {
	return OrphanDiskCmd{ui: ui, director: director}
}

func (c OrphanDiskCmd) Run(opts OrphanDiskOpts) error {
	err := c.ui.AskForConfirmation()
	if err != nil {
		return err
	}

	return c.director.OrphanDisk(opts.Args.CID)
}
