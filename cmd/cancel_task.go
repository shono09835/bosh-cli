package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
)

type CancelTaskCmd struct {
	director boshdir.Director
}

func NewCancelTaskCmd(director boshdir.Director) CancelTaskCmd {
	return CancelTaskCmd{director: director}
}

func (c CancelTaskCmd) Run(opts CancelTaskOpts) error {
	task, err := c.director.FindTask(opts.Args.ID)
	if err != nil {
		return err
	}

	return task.Cancel()
}
