package stemcell

import (
	bideplmanifest "github.com/shono09835/bosh-cli/v7/deployment/manifest"
	bitarball "github.com/shono09835/bosh-cli/v7/installation/tarball"
	biui "github.com/shono09835/bosh-cli/v7/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Fetcher struct {
	TarballProvider   bitarball.Provider
	StemcellExtractor Extractor
}

func (s Fetcher) GetStemcell(deploymentManifest bideplmanifest.Manifest, stage biui.Stage) (ExtractedStemcell, error) {
	stemcell, err := deploymentManifest.Stemcell(deploymentManifest.JobName())
	if err != nil {
		return nil, err
	}

	stemcellTarballPath, err := s.TarballProvider.Get(stemcell, stage)
	if err != nil {
		return nil, err
	}

	var extractedStemcell ExtractedStemcell
	err = stage.Perform("Validating stemcell", func() error {
		extractedStemcell, err = s.StemcellExtractor.Extract(stemcellTarballPath)
		if err != nil {
			return bosherr.WrapErrorf(err, "Extracting stemcell from '%s'", stemcellTarballPath)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return extractedStemcell, nil
}
