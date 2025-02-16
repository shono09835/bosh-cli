package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshreldir "github.com/shono09835/bosh-cli/v7/releasedir"
	biui "github.com/shono09835/bosh-cli/v7/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type RemoveBlobCmd struct {
	blobsDir boshreldir.BlobsDir
	ui       biui.UI
}

func NewRemoveBlobCmd(blobsDir boshreldir.BlobsDir, ui biui.UI) RemoveBlobCmd {
	return RemoveBlobCmd{blobsDir: blobsDir, ui: ui}
}

func (c RemoveBlobCmd) Run(opts RemoveBlobOpts) error {
	err := c.blobsDir.UntrackBlob(opts.Args.BlobsPath)
	if err != nil {
		return bosherr.WrapErrorf(err, "Untracking blob")
	}

	c.ui.PrintLinef("Removed blob '%s'", opts.Args.BlobsPath)

	return nil
}
