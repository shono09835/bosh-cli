package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
)

type DiffConfigCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewDiffConfigCmd(ui boshui.UI, director boshdir.Director) DiffConfigCmd {
	return DiffConfigCmd{ui: ui, director: director}
}

func (c DiffConfigCmd) Run(opts DiffConfigOpts) error {
	configDiff, err := c.director.DiffConfigByIDOrContent(opts.FromID, opts.FromContent.Bytes, opts.ToID, opts.ToContent.Bytes)
	if err != nil {
		return err
	}

	diff := NewDiff(configDiff.Diff)

	ConfigDiffTable{diff, opts, c.ui}.Print()

	return nil
}
