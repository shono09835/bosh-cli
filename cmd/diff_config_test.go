package cmd_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/cmd"
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	fakedir "github.com/shono09835/bosh-cli/v7/director/directorfakes"
	fakeui "github.com/shono09835/bosh-cli/v7/ui/fakes"
	boshtbl "github.com/shono09835/bosh-cli/v7/ui/table"
)

var _ = Describe("DiffConfigCmd", func() {
	var (
		ui       *fakeui.FakeUI
		director *fakedir.FakeDirector
		command  DiffConfigCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		director = &fakedir.FakeDirector{}
		command = NewDiffConfigCmd(ui, director)
	})

	Describe("Run", func() {
		var (
			opts DiffConfigOpts
		)

		BeforeEach(func() {
			opts = DiffConfigOpts{
				FromID: "1",
				ToID:   "2",
			}
		})

		act := func() error { return command.Run(opts) }

		It("diff two configs", func() {
			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(director.DiffConfigByIDOrContentCallCount()).To(Equal(1))

			from, _, to, _ := director.DiffConfigByIDOrContentArgsForCall(0)
			Expect(from).To(Equal("1"))
			Expect(to).To(Equal("2"))
		})

		It("returns an error if diffing failed", func() {
			director.DiffConfigByIDOrContentReturns(boshdir.ConfigDiff{}, errors.New("Fetching diff result"))

			err := act()
			Expect(err).To(HaveOccurred())
		})

		It("gets the diff from the two configs", func() {
			diff := [][]interface{}{
				{"some line that stayed", nil},
				{"some line that was added", "added"},
				{"some line that was removed", "removed"},
			}

			expectedDiff := boshdir.NewConfigDiff(diff)
			director.DiffConfigByIDOrContentReturns(expectedDiff, nil)
			err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(ui.Table).To(Equal(
				boshtbl.Table{
					Content: "",

					Header: []boshtbl.Header{
						boshtbl.NewHeader("From ID"),
						boshtbl.NewHeader("To ID"),
						boshtbl.NewHeader("Diff"),
					},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("1"),
							boshtbl.NewValueString("2"),
							boshtbl.NewValueString("  some line that stayed\n+ some line that was added\n- some line that was removed\n"),
						},
					},

					Notes: []string{},

					FillFirstColumn: true,

					Transpose: true,
				}))
			Expect(director.DiffConfigByIDOrContentCallCount()).To(Equal(1))
		})
	})
})
