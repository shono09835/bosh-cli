package state

import (
	biagentclient "github.com/cloudfoundry/bosh-agent/agentclient"
	biblobstore "github.com/shono09835/bosh-cli/v7/blobstore"
	bideplrel "github.com/shono09835/bosh-cli/v7/deployment/release"
	bistatejob "github.com/shono09835/bosh-cli/v7/state/job"
	bistatepkg "github.com/shono09835/bosh-cli/v7/state/pkg"
	bitemplate "github.com/shono09835/bosh-cli/v7/templatescompiler"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type BuilderFactory interface {
	NewBuilder(biblobstore.Blobstore, biagentclient.AgentClient) Builder
}

type builderFactory struct {
	packageRepo               bistatepkg.CompiledPackageRepo
	releaseJobResolver        bideplrel.JobResolver
	jobRenderer               bitemplate.JobListRenderer
	renderedJobListCompressor bitemplate.RenderedJobListCompressor
	logger                    boshlog.Logger
}

func NewBuilderFactory(
	packageRepo bistatepkg.CompiledPackageRepo,
	releaseJobResolver bideplrel.JobResolver,
	jobRenderer bitemplate.JobListRenderer,
	renderedJobListCompressor bitemplate.RenderedJobListCompressor,
	logger boshlog.Logger,
) BuilderFactory {
	return &builderFactory{
		packageRepo:               packageRepo,
		releaseJobResolver:        releaseJobResolver,
		jobRenderer:               jobRenderer,
		renderedJobListCompressor: renderedJobListCompressor,
		logger:                    logger,
	}
}

func (f *builderFactory) NewBuilder(blobstore biblobstore.Blobstore, agentClient biagentclient.AgentClient) Builder {
	packageCompiler := NewRemotePackageCompiler(blobstore, agentClient, f.packageRepo)
	jobDependencyCompiler := bistatejob.NewDependencyCompiler(packageCompiler, f.logger)

	return NewBuilder(
		f.releaseJobResolver,
		jobDependencyCompiler,
		f.jobRenderer,
		f.renderedJobListCompressor,
		blobstore,
		f.logger,
	)
}
