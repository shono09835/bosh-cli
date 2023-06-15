package cmd

import (
	"github.com/cppforlife/go-patch/patch"

	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshtpl "github.com/shono09835/bosh-cli/v7/director/template"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
)

type DeleteEnvCmd struct {
	ui          boshui.UI
	envProvider func(string, string, boshtpl.Variables, patch.Op) DeploymentDeleter
}

func NewDeleteEnvCmd(ui boshui.UI, envProvider func(string, string, boshtpl.Variables, patch.Op) DeploymentDeleter) *DeleteEnvCmd {
	return &DeleteEnvCmd{ui: ui, envProvider: envProvider}
}

func (c *DeleteEnvCmd) Run(stage boshui.Stage, opts DeleteEnvOpts) error {
	c.ui.BeginLinef("Deployment manifest: '%s'\n", opts.Args.Manifest.Path)

	depDeleter := c.envProvider(
		opts.Args.Manifest.Path, opts.StatePath, opts.VarFlags.AsVariables(), opts.OpsFlags.AsOp())

	return depDeleter.DeleteDeployment(opts.SkipDrain, stage)
}
