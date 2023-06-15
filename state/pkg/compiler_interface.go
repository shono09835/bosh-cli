package pkg

import (
	birelpkg "github.com/shono09835/bosh-cli/v7/release/pkg"
)

type Compiler interface {
	Compile(birelpkg.Compilable) (CompiledPackageRecord, bool, error)
}
