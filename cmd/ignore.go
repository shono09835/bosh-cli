package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
)

type IgnoreCmd struct {
	deployment boshdir.Deployment
}

func NewIgnoreCmd(deployment boshdir.Deployment) IgnoreCmd {
	return IgnoreCmd{deployment: deployment}
}

func (cmd IgnoreCmd) Run(opts IgnoreOpts) error {
	return cmd.deployment.Ignore(opts.Args.Slug, true)
}
