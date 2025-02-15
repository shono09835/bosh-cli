package cmd_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/cmd"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
)

var _ = Describe("InstanceTable", func() {
	Describe("ForVMInfo", func() {
		var (
			info boshdir.VMInfo
			tbl  InstanceTable
		)

		BeforeEach(func() {
			info = boshdir.VMInfo{}
			tbl = InstanceTable{Details: true, DNS: true, Vitals: true, CloudProperties: true}
		})

		Describe("name, id", func() {
			It("returns ? name", func() {
				Expect(tbl.ForVMInfo(info).Name.String()).To(Equal("?"))
			})

			It("returns name", func() {
				info.JobName = "name"
				Expect(tbl.ForVMInfo(info).Name.String()).To(Equal("name"))
			})

			It("returns name with id", func() {
				info.JobName = "name"
				info.ID = "id"
				Expect(tbl.ForVMInfo(info).Name.String()).To(Equal("name/id"))
			})

			It("returns name with id, bootstrap and index", func() {
				idx := 1
				info.JobName = "name"
				info.ID = "id"
				info.Index = &idx
				info.Bootstrap = true
				Expect(tbl.ForVMInfo(info).Name.String()).To(Equal("name/id"))

				Expect(tbl.ForVMInfo(info).Bootstrap).ToNot(BeNil())
				Expect(tbl.ForVMInfo(info).Bootstrap.String()).To(Equal("true"))
				Expect(tbl.ForVMInfo(info).Index).ToNot(BeNil())
				Expect(tbl.ForVMInfo(info).Index.String()).To(Equal("1"))
			})

			It("returns name with id, without index", func() {
				info.JobName = "name"
				info.ID = "id"
				Expect(tbl.ForVMInfo(info).Name.String()).To(Equal("name/id"))
			})

			It("returns ? name with id", func() {
				info.JobName = ""
				info.ID = "id"
				Expect(tbl.ForVMInfo(info).Name.String()).To(Equal("?/id"))
			})
		})

		Describe("vm type, resource pool", func() {
			It("returns RP if vm type is empty", func() {
				info.ResourcePool = "rp"
				Expect(tbl.ForVMInfo(info).VMType.String()).To(Equal("rp"))
			})

			It("returns vm type if vm type is non-empty", func() {
				info.ResourcePool = "rp"
				info.VMType = "vm-type"
				Expect(tbl.ForVMInfo(info).VMType.String()).To(Equal("vm-type"))
			})
		})

		Describe("vm created at", func() {
			It("returns empty if vm_created_at is empty", func() {
				info.VMCreatedAt = time.Time{}
				Expect(tbl.ForVMInfo(info).VMCreatedAt.String()).To(Equal(""))
				info.VMCreatedAt = time.Unix(0, 0).UTC()
				Expect(tbl.ForVMInfo(info).VMCreatedAt.String()).To(Equal(""))
			})

			It("returns time if created_at is non-empty", func() {
				info.VMCreatedAt = time.Date(2016, time.January, 9, 6, 23, 25, 0, time.UTC)
				Expect(tbl.ForVMInfo(info).VMCreatedAt.String()).To(Equal("Sat Jan  9 06:23:25 UTC 2016"))
			})
		})

		Describe("disk cids", func() {
			It("returns empty if disk cids is empty", func() {
				Expect(tbl.ForVMInfo(info).DiskCIDs.String()).To(Equal(""))
			})

			It("returns disk cid if disk cids is non-empty", func() {
				info.DiskIDs = []string{"disk-cid1", "disk-cid2"}
				Expect(tbl.ForVMInfo(info).DiskCIDs.String()).To(Equal("disk-cid1\ndisk-cid2"))
			})
		})
	})
})
