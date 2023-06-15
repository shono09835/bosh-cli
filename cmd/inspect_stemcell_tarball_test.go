package cmd_test

import (
	"errors"

	. "github.com/shono09835/bosh-cli/v7/cmd"
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	fakedir "github.com/shono09835/bosh-cli/v7/director/directorfakes"
	fakeui "github.com/shono09835/bosh-cli/v7/ui/fakes"
	boshtbl "github.com/shono09835/bosh-cli/v7/ui/table"
	biproperty "github.com/cloudfoundry/bosh-utils/property"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InspectStemcellTarballCmd", func() {
	Describe("Run", func() {
		var (
			fs               *fakesys.FakeFileSystem
			archive          *fakedir.FakeStemcellArchive
			command          InspectStemcellTarballCmd
			ui               *fakeui.FakeUI
			opts             InspectStemcellTarballOpts
			stemcellMetadata boshdir.StemcellMetadata
		)

		BeforeEach(func() {
			fs = fakesys.NewFakeFileSystem()
			archive = &fakedir.FakeStemcellArchive{}
			stemcellMetadata = boshdir.StemcellMetadata{
				Name:    "example-name",
				OS:      "example-os",
				Version: "example.version",
				CloudProperties: biproperty.Map{
					"infrastructure": "example-infrastructure",
					"hypervisor":     "example-hypervisor",
				},
			}

			stemcellArchiveFactory := func(path string) boshdir.StemcellArchive {
				if archive.FileStub == nil {
					archive.FileStub = func() (boshdir.UploadFile, error) {
						return fakesys.NewFakeFile(path, fs), nil
					}
				}
				return archive
			}

			opts = InspectStemcellTarballOpts{}
			ui = &fakeui.FakeUI{}

			command = NewInspectStemcellTarballCmd(stemcellArchiveFactory, ui)
		})

		Context("when infrastructure is known", func() {
			It("returns a table with name, os, version, and infrastructure", func() {
				archive.InfoReturns(stemcellMetadata, nil)

				err := command.Run(opts)
				Expect(err).ToNot(HaveOccurred())

				Expect(ui.Table).To(Equal(boshtbl.Table{
					Content: "stemcell-metadata",

					Header: []boshtbl.Header{
						boshtbl.NewHeader("Name"),
						boshtbl.NewHeader("OS"),
						boshtbl.NewHeader("Version"),
						boshtbl.NewHeader("Infrastructure"),
						boshtbl.NewHeader("Hypervisor"),
					},

					SortBy: []boshtbl.ColumnSort{
						{Column: 0, Asc: true},
					},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("example-name"),
							boshtbl.NewValueString("example-os"),
							boshtbl.NewValueString("example.version"),
							boshtbl.NewValueString("example-infrastructure"),
							boshtbl.NewValueString("example-hypervisor"),
						},
					},
				}))
			})
		})

		Context("when infrastructure is unknown", func() {

			BeforeEach(func() {
				stemcellMetadata.CloudProperties["infrastructure"] = nil
			})

			It("returns a table with name, os, version, and infrastructure", func() {
				archive.InfoReturns(stemcellMetadata, nil)

				err := command.Run(opts)
				Expect(err).ToNot(HaveOccurred())

				Expect(ui.Table).To(Equal(boshtbl.Table{
					Content: "stemcell-metadata",

					Header: []boshtbl.Header{
						boshtbl.NewHeader("Name"),
						boshtbl.NewHeader("OS"),
						boshtbl.NewHeader("Version"),
						boshtbl.NewHeader("Infrastructure"),
						boshtbl.NewHeader("Hypervisor"),
					},

					SortBy: []boshtbl.ColumnSort{
						{Column: 0, Asc: true},
					},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("example-name"),
							boshtbl.NewValueString("example-os"),
							boshtbl.NewValueString("example.version"),
							boshtbl.NewValueString("unknown"),
							boshtbl.NewValueString("example-hypervisor"),
						},
					},
				}))
			})
		})

		Context("when hypervisor is unknown", func() {

			BeforeEach(func() {
				stemcellMetadata.CloudProperties["hypervisor"] = nil
			})

			It("returns a table with name, os, version, and infrastructure", func() {
				archive.InfoReturns(stemcellMetadata, nil)

				err := command.Run(opts)
				Expect(err).ToNot(HaveOccurred())

				Expect(ui.Table).To(Equal(boshtbl.Table{
					Content: "stemcell-metadata",

					Header: []boshtbl.Header{
						boshtbl.NewHeader("Name"),
						boshtbl.NewHeader("OS"),
						boshtbl.NewHeader("Version"),
						boshtbl.NewHeader("Infrastructure"),
						boshtbl.NewHeader("Hypervisor"),
					},

					SortBy: []boshtbl.ColumnSort{
						{Column: 0, Asc: true},
					},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("example-name"),
							boshtbl.NewValueString("example-os"),
							boshtbl.NewValueString("example.version"),
							boshtbl.NewValueString("example-infrastructure"),
							boshtbl.NewValueString("-"),
						},
					},
				}))
			})
		})

		It("returns error if retrieving stemcell archive info fails", func() {
			archive.InfoReturns(boshdir.StemcellMetadata{}, errors.New("fake-err"))

			err := command.Run(opts)
			Expect(err).To(HaveOccurred())
		})
	})
})
