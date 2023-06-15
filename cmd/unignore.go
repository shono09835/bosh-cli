package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
)

type UnignoreCmd struct {
	deployment boshdir.Deployment
}

func NewUnignoreCmd(deployment boshdir.Deployment) UnignoreCmd {
	return UnignoreCmd{deployment: deployment}
}

func (cmd UnignoreCmd) Run(opts UnignoreOpts) error {
	return cmd.deployment.Ignore(opts.Args.Slug, false)
}
