package cmd_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/cmd"
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	fakereldir "github.com/shono09835/bosh-cli/v7/releasedir/releasedirfakes"
)

var _ = Describe("GenerateJobCmd", func() {
	var (
		releaseDir *fakereldir.FakeReleaseDir
		command    GenerateJobCmd
	)

	BeforeEach(func() {
		releaseDir = &fakereldir.FakeReleaseDir{}
		command = NewGenerateJobCmd(releaseDir)
	})

	Describe("Run", func() {
		var (
			opts GenerateJobOpts
		)

		BeforeEach(func() {
			opts = GenerateJobOpts{Args: GenerateJobArgs{Name: "job"}}
		})

		act := func() error { return command.Run(opts) }

		It("generates job", func() {
			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(releaseDir.GenerateJobCallCount()).To(Equal(1))
			Expect(releaseDir.GenerateJobArgsForCall(0)).To(Equal("job"))
		})

		It("returns error if generating job fails", func() {
			releaseDir.GenerateJobReturns(errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
