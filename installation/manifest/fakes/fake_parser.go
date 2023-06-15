package fakes

import (
	"github.com/cppforlife/go-patch/patch"

	boshtpl "github.com/shono09835/bosh-cli/v7/director/template"
	biinstallmanifest "github.com/shono09835/bosh-cli/v7/installation/manifest"
	birelsetmanifest "github.com/shono09835/bosh-cli/v7/release/set/manifest"
)

type FakeParser struct {
	ParsePath          string
	ReleaseSetManifest birelsetmanifest.Manifest
	ParseManifest      biinstallmanifest.Manifest
	ParseErr           error
}

func NewFakeParser() *FakeParser {
	return &FakeParser{}
}

func (p *FakeParser) Parse(path string, vars boshtpl.Variables, op patch.Op, releaseSetManifest birelsetmanifest.Manifest) (biinstallmanifest.Manifest, error) {
	p.ParsePath = path
	p.ReleaseSetManifest = releaseSetManifest
	return p.ParseManifest, p.ParseErr
}
