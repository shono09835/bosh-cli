package opts_test

import (
	"errors"

	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	"github.com/cppforlife/go-patch/patch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OpsFileArg", func() {
	Describe("UnmarshalFlag", func() {
		var (
			fs  *fakesys.FakeFileSystem
			arg OpsFileArg
		)

		BeforeEach(func() {
			fs = fakesys.NewFakeFileSystem()
			arg = OpsFileArg{FS: fs}
		})

		It("sets read operations", func() {
			err := fs.WriteFileString("/some/path", `
- type: remove
  path: /a
- type: remove
  path: /b
`)
			Expect(err).ToNot(HaveOccurred())

			err = (&arg).UnmarshalFlag("/some/path")
			Expect(err).ToNot(HaveOccurred())

			Expect(arg.Ops).To(Equal(patch.Ops{
				patch.DescriptiveOp{
					Op: patch.RemoveOp{
						Path: patch.MustNewPointerFromString("/a"),
					},
					ErrorMsg: "operation [0] in /some/path failed",
				},
				patch.DescriptiveOp{
					Op: patch.RemoveOp{
						Path: patch.MustNewPointerFromString("/b"),
					},
					ErrorMsg: "operation [1] in /some/path failed",
				},
			}))
		})

		It("returns an error if operations are not valid", func() {
			err := fs.WriteFileString("/some/path", "- type: unknown")
			Expect(err).ToNot(HaveOccurred())

			err = (&arg).UnmarshalFlag("/some/path")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(`Building ops: Unknown operation [0] with type 'unknown' within
{
  "Type": "unknown",
  "Error": "operation [0] in /some/path failed"
}`))
		})

		It("returns an error if reading file fails", func() {
			err := fs.WriteFileString("/some/path", "content")
			Expect(err).ToNot(HaveOccurred())
			fs.ReadFileError = errors.New("fake-err")

			err = (&arg).UnmarshalFlag("/some/path")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("returns an error if parsing file fails", func() {
			err := fs.WriteFileString("/some/path", "content")
			Expect(err).ToNot(HaveOccurred())

			err = (&arg).UnmarshalFlag("/some/path")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Deserializing ops file '/some/path'"))
		})

		It("returns an error when it's empty", func() {
			err := (&arg).UnmarshalFlag("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected file path to be non-empty"))
		})
	})
})
