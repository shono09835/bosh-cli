package opts_test

import (
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DirOrCWDArg", func() {
	Describe("UnmarshalFlag", func() {
		var (
			arg DirOrCWDArg
		)

		BeforeEach(func() {
			arg = DirOrCWDArg{}
		})

		It("returns with path set", func() {
			err := (&arg).UnmarshalFlag("/some/path")
			Expect(err).ToNot(HaveOccurred())
			Expect(arg.Path).To(Equal("/some/path"))
		})

		It("returns with cwd path set when it's empty", func() {
			err := (&arg).UnmarshalFlag("")
			Expect(err).ToNot(HaveOccurred())
			Expect(arg.Path).ToNot(BeEmpty())
		})

		It("returns with cwd path set when it's '.'", func() {
			err := (&arg).UnmarshalFlag(".")
			Expect(err).ToNot(HaveOccurred())
			Expect(arg.Path).ToNot(BeEmpty())
		})
	})
})
