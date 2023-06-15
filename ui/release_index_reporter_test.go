package ui_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/ui"
	fakeui "github.com/shono09835/bosh-cli/v7/ui/fakes"
)

var _ = Describe("ReleaseIndexReporter", func() {
	var (
		ui       *fakeui.FakeUI
		reporter ReleaseIndexReporter
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		reporter = NewReleaseIndexReporter(ui)
	})

	Describe("ReleaseIndexAdded", func() {
		It("prints failed msg", func() {
			reporter.ReleaseIndexAdded("name", "desc", errors.New("err"))
			Expect(ui.Errors).To(Equal([]string{"Failed adding name release 'desc'"}))
		})

		It("prints finished msg", func() {
			reporter.ReleaseIndexAdded("name", "desc", nil)
			Expect(ui.Said).To(Equal([]string{"Added name release 'desc'"}))
		})
	})
})
