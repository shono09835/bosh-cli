package cmd

import (
	"github.com/cppforlife/go-patch/patch"

	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshtpl "github.com/shono09835/bosh-cli/v7/director/template"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
)

type CreateEnvCmd struct {
	ui          boshui.UI
	envProvider EnvProviderFunction
}

type EnvProviderFunction func(string, string, boshtpl.Variables, patch.Op) DeploymentPreparer

func NewCreateEnvCmd(ui boshui.UI, envProvider EnvProviderFunction) *CreateEnvCmd {
	return &CreateEnvCmd{ui: ui, envProvider: envProvider}
}

func (c *CreateEnvCmd) Run(stage boshui.Stage, opts CreateEnvOpts) error {
	c.ui.BeginLinef("Deployment manifest: '%s'\n", opts.Args.Manifest.Path)

	depPreparer := c.envProvider(opts.Args.Manifest.Path, opts.StatePath, opts.VarFlags.AsVariables(), opts.OpsFlags.AsOp())

	return depPreparer.PrepareDeployment(stage, opts.Recreate, opts.RecreatePersistentDisks, opts.SkipDrain)
}
