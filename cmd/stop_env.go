package cmd

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshtpl "github.com/shono09835/bosh-cli/v7/director/template"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
	"github.com/cppforlife/go-patch/patch"
)

type StopEnvCmd struct {
	ui          boshui.UI
	envProvider func(string, string, boshtpl.Variables, patch.Op) DeploymentStateManager
}

func NewStopEnvCmd(ui boshui.UI, envProvider func(string, string, boshtpl.Variables, patch.Op) DeploymentStateManager) *StopEnvCmd {
	return &StopEnvCmd{ui: ui, envProvider: envProvider}
}

func (c *StopEnvCmd) Run(stage boshui.Stage, opts StopEnvOpts) error {
	c.ui.BeginLinef("Deployment manifest: '%s'\n", opts.Args.Manifest.Path)

	depStateManager := c.envProvider(
		opts.Args.Manifest.Path, opts.StatePath, opts.VarFlags.AsVariables(), opts.OpsFlags.AsOp())

	return depStateManager.StopDeployment(opts.SkipDrain, stage)
}
