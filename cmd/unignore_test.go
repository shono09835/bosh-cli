package cmd_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/cmd"
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"

	"errors"

	boshdir "github.com/shono09835/bosh-cli/v7/director"
	fakedir "github.com/shono09835/bosh-cli/v7/director/directorfakes"
)

var _ = Describe("UnignoreCmd", func() {
	var (
		deployment *fakedir.FakeDeployment
		command    UnignoreCmd
	)

	BeforeEach(func() {
		deployment = &fakedir.FakeDeployment{}
		command = NewUnignoreCmd(deployment)
	})

	Describe("Run", func() {
		var (
			opts UnignoreOpts
		)

		BeforeEach(func() {
			opts = UnignoreOpts{}
		})

		act := func() error {
			return command.Run(opts)
		}

		Context("when unignoring an instance", func() {
			BeforeEach(func() {
				opts.Args.Slug = boshdir.NewInstanceSlug("some-name", "some-id")
			})

			It("unignores the instance", func() {
				err := act()
				Expect(err).ToNot(HaveOccurred())

				Expect(deployment.IgnoreCallCount()).To(Equal(1))

				slugArg, unignoreArg := deployment.IgnoreArgsForCall(0)
				Expect(slugArg).To(Equal(boshdir.NewInstanceSlug("some-name", "some-id")))
				Expect(unignoreArg).To(Equal(false))
			})

			Context("when unignoring fails", func() {

				BeforeEach(func() {
					deployment.IgnoreReturns(errors.New("nope nope nope"))
				})

				It("returns the error", func() {
					err := act()
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
