package stemcell_test

import (
	"errors"

	. "github.com/shono09835/bosh-cli/v7/stemcell"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	fakecmd "github.com/cloudfoundry/bosh-utils/fileutil/fakes"
	biproperty "github.com/cloudfoundry/bosh-utils/property"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
)

var _ = Describe("Reader", func() {
	var (
		compressor     *fakecmd.FakeCompressor
		stemcellReader Reader
		fs             *fakesys.FakeFileSystem
	)

	BeforeEach(func() {
		compressor = fakecmd.NewFakeCompressor()
		fs = fakesys.NewFakeFileSystem()
		stemcellReader = NewReader(compressor, fs)

		manifestContents := `
---
name: fake-stemcell-name
version: '2690'
operating_system: ubuntu-trusty
sha1: sha
bosh_protocol: 1
stemcell_formats: ['aws-raw']
cloud_properties:
  infrastructure: aws
  ami:
    us-east-1: fake-ami-version
    `
		err := fs.WriteFileString("fake-extracted-path/stemcell.MF", manifestContents)
		Expect(err).ToNot(HaveOccurred())
	})

	It("extracts the stemcells from a stemcell path", func() {
		_, err := stemcellReader.Read("fake-stemcell-path", "fake-extracted-path")
		Expect(err).ToNot(HaveOccurred())
		Expect(compressor.DecompressFileToDirTarballPaths).To(ContainElement("fake-stemcell-path"))
		Expect(compressor.DecompressFileToDirDirs).To(ContainElement("fake-extracted-path"))
	})

	It("generates correct stemcell", func() {
		stemcell, err := stemcellReader.Read("fake-stemcell-path", "fake-extracted-path")
		Expect(err).ToNot(HaveOccurred())
		expectedStemcell := NewExtractedStemcell(
			Manifest{
				Name:            "fake-stemcell-name",
				Version:         "2690",
				OS:              "ubuntu-trusty",
				SHA1:            "sha",
				BoshProtocol:    "1",
				StemcellFormats: []string{"aws-raw"},
				CloudProperties: biproperty.Map{
					"infrastructure": "aws",
					"ami": biproperty.Map{
						"us-east-1": "fake-ami-version",
					},
				},
			},
			"fake-extracted-path",
			compressor,
			fs,
		)
		Expect(stemcell.Manifest().CloudProperties).To(Equal(expectedStemcell.Manifest().CloudProperties))
		Expect(stemcell).To(Equal(expectedStemcell))

	})

	Context("when api_version is specified", func() {
		BeforeEach(func() {
			manifestContents := `
---
name: fake-stemcell-name
version: '2690'
operating_system: ubuntu-trusty
sha1: sha
bosh_protocol: 1
stemcell_formats: ['aws-raw']
api_version: 2
cloud_properties:
  infrastructure: aws
  ami:
    us-east-1: fake-ami-version
    `
			err := fs.WriteFileString("fake-extracted-path/stemcell.MF", manifestContents)
			Expect(err).ToNot(HaveOccurred())
		})

		It("generates correct stemcell", func() {
			stemcell, err := stemcellReader.Read("fake-stemcell-path", "fake-extracted-path")
			Expect(err).ToNot(HaveOccurred())
			expectedStemcell := NewExtractedStemcell(
				Manifest{
					Name:            "fake-stemcell-name",
					Version:         "2690",
					OS:              "ubuntu-trusty",
					SHA1:            "sha",
					BoshProtocol:    "1",
					ApiVersion:      2,
					StemcellFormats: []string{"aws-raw"},
					CloudProperties: biproperty.Map{
						"infrastructure": "aws",
						"ami": biproperty.Map{
							"us-east-1": "fake-ami-version",
						},
					},
				},
				"fake-extracted-path",
				compressor,
				fs,
			)
			Expect(stemcell.Manifest().CloudProperties).To(Equal(expectedStemcell.Manifest().CloudProperties))
			Expect(stemcell).To(Equal(expectedStemcell))

		})

	})

	Context("when extracting stemcell fails", func() {
		BeforeEach(func() {
			compressor.DecompressFileToDirErr = errors.New("fake-decompress-error")
		})

		It("returns an error", func() {
			_, err := stemcellReader.Read("fake-stemcell-path", "fake-extracted-path")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-decompress-error"))
		})
	})

	Context("when reading stemcell manifest fails", func() {
		BeforeEach(func() {
			fs.ReadFileError = errors.New("fake-read-error")
		})

		It("returns an error", func() {
			_, err := stemcellReader.Read("fake-stemcell-path", "fake-extracted-path")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-read-error"))
		})
	})

	Context("when parsing stemcell manifest fails", func() {
		BeforeEach(func() {
			err := fs.WriteFileString("fake-extracted-path/stemcell.MF", "<not-a-yaml>")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error", func() {
			_, err := stemcellReader.Read("fake-stemcell-path", "fake-extracted-path")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing stemcell manifest"))
		})
	})

})
