package cmd_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/cmd"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	fakedir "github.com/shono09835/bosh-cli/v7/director/directorfakes"
	fakeui "github.com/shono09835/bosh-cli/v7/ui/fakes"
	boshtbl "github.com/shono09835/bosh-cli/v7/ui/table"
)

var _ = Describe("LocksCmd", func() {
	var (
		ui       *fakeui.FakeUI
		director *fakedir.FakeDirector
		command  LocksCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		director = &fakedir.FakeDirector{}
		command = NewLocksCmd(ui, director)
	})

	Describe("Run", func() {
		act := func() error { return command.Run() }

		It("lists current locks", func() {
			locks := []boshdir.Lock{
				{
					Type:      "deployment",
					Resource:  []string{"some-deployment", "20"},
					ExpiresAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					TaskID:    "123456",
				},
			}

			director.LocksReturns(locks, nil)

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(ui.Table).To(Equal(boshtbl.Table{
				Content: "locks",

				Header: []boshtbl.Header{
					boshtbl.NewHeader("Type"),
					boshtbl.NewHeader("Resource"),
					boshtbl.NewHeader("Task ID"),
					boshtbl.NewHeader("Expires at"),
				},

				SortBy: []boshtbl.ColumnSort{{Column: 2, Asc: true}},

				Rows: [][]boshtbl.Value{
					{
						boshtbl.NewValueString("deployment"),
						boshtbl.NewValueString("some-deployment:20"),
						boshtbl.NewValueString("123456"),
						boshtbl.NewValueTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
					},
				},
			}))
		})

		It("returns error if locks cannot be retrieved", func() {
			director.LocksReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
