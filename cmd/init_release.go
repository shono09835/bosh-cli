package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshreldir "github.com/shono09835/bosh-cli/v7/releasedir"
)

type InitReleaseCmd struct {
	releaseDir boshreldir.ReleaseDir
}

func NewInitReleaseCmd(releaseDir boshreldir.ReleaseDir) InitReleaseCmd {
	return InitReleaseCmd{releaseDir: releaseDir}
}

func (c InitReleaseCmd) Run(opts InitReleaseOpts) error {
	return c.releaseDir.Init(opts.Git)
}
