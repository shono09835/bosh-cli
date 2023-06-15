package fakes

import (
	"github.com/cppforlife/go-patch/patch"

	boshtpl "github.com/shono09835/bosh-cli/v7/director/template"
	birelsetmanifest "github.com/shono09835/bosh-cli/v7/release/set/manifest"
)

type FakeParser struct {
	ParsePath     string
	ParseManifest birelsetmanifest.Manifest
	ParseErr      error
}

func NewFakeParser() *FakeParser {
	return &FakeParser{}
}

func (p *FakeParser) Parse(path string, vars boshtpl.Variables, op patch.Op) (birelsetmanifest.Manifest, error) {
	p.ParsePath = path
	return p.ParseManifest, p.ParseErr
}
