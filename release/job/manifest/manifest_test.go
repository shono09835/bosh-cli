package manifest_test

import (
	"errors"

	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/release/job/manifest"
)

var _ = Describe("NewManifestFromPath", func() {
	var (
		fs *fakesys.FakeFileSystem
	)

	BeforeEach(func() {
		fs = fakesys.NewFakeFileSystem()
	})

	It("parses job manifest successfully", func() {
		contents := `---
name: name

templates:
  src.yml: dst.yml

packages:
- pkg1
- pkg2

properties:
  prop1:
    description: prop1-desc
    default: 1
  prop1.prop2:
    description: prop2-desc
    default: prop2-default
`

		err := fs.WriteFileString("/path", contents)
		Expect(err).ToNot(HaveOccurred())

		manifest, err := NewManifestFromPath("/path", fs)
		Expect(err).ToNot(HaveOccurred())
		Expect(manifest).To(Equal(Manifest{
			Name: "name",

			Templates: map[string]string{"src.yml": "dst.yml"},

			Packages: []string{"pkg1", "pkg2"},

			Properties: map[string]PropertyDefinition{
				"prop1": PropertyDefinition{
					Description: "prop1-desc",
					Default:     1,
				},
				"prop1.prop2": PropertyDefinition{
					Description: "prop2-desc",
					Default:     "prop2-default",
				},
			},
		}))
	})

	It("returns error if manifest is not valid yaml", func() {
		err := fs.WriteFileString("/path", "-")
		Expect(err).ToNot(HaveOccurred())

		_, err = NewManifestFromPath("/path", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("line 1"))
	})

	It("returns error if manifest cannot be read", func() {
		err := fs.WriteFileString("/path", "-")
		Expect(err).ToNot(HaveOccurred())
		fs.ReadFileError = errors.New("fake-err")

		_, err = NewManifestFromPath("/path", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("fake-err"))
	})
})
