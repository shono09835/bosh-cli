package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
)

type CancelTasksCmd struct {
	director boshdir.Director
}

func NewCancelTasksCmd(director boshdir.Director) CancelTasksCmd {
	return CancelTasksCmd{director: director}
}

func (c CancelTasksCmd) Run(opts CancelTasksOpts) error {
	filter := boshdir.TasksFilter{
		Deployment: opts.Deployment,
		Types:      opts.Types,
		States:     opts.States,
	}

	return c.director.CancelTasks(filter)
}
