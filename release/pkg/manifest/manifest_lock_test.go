package manifest_test

import (
	"errors"

	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/release/pkg/manifest"
)

var _ = Describe("NewManifestLockFromPath", func() {
	var (
		fs *fakesys.FakeFileSystem
	)

	BeforeEach(func() {
		fs = fakesys.NewFakeFileSystem()
	})

	It("parses pkg manifest successfully", func() {
		contents := `---
name: name
fingerprint: fp
dependencies:
- pkg1
- pkg2
`

		err := fs.WriteFileString("/path", contents)
		Expect(err).ToNot(HaveOccurred())

		manifest, err := NewManifestLockFromPath("/path", fs)
		Expect(err).ToNot(HaveOccurred())
		Expect(manifest).To(Equal(ManifestLock{
			Name:         "name",
			Fingerprint:  "fp",
			Dependencies: []string{"pkg1", "pkg2"},
		}))
	})

	It("returns error if manifest is not valid yaml", func() {
		err := fs.WriteFileString("/path", "-")
		Expect(err).ToNot(HaveOccurred())

		_, err = NewManifestLockFromPath("/path", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("line 1"))
	})

	It("returns error if manifest cannot be read", func() {
		err := fs.WriteFileString("/path", "-")
		Expect(err).ToNot(HaveOccurred())
		fs.ReadFileError = errors.New("fake-err")

		_, err = NewManifestLockFromPath("/path", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("fake-err"))
	})
})

var _ = Describe("ManifestLock", func() {
	Describe("AsBytes", func() {
		It("returns serializes manifest", func() {
			bytes, err := ManifestLock{
				Name: "name", Fingerprint: "fp", Dependencies: []string{"pkg1", "pkg2"}}.AsBytes()
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(`name: name
fingerprint: fp
dependencies:
- pkg1
- pkg2
`))
		})

		It("returns error if name is empty", func() {
			_, err := ManifestLock{}.AsBytes()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected non-empty package name"))
		})

		It("returns error if fingerprint is empty", func() {
			_, err := ManifestLock{Name: "name"}.AsBytes()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected non-empty package fingerprint"))
		})
	})
})
