package templatescompiler_test

import (
	"bytes"
	"os"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	fakeboshsys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	boshreljob "github.com/shono09835/bosh-cli/v7/release/job"
	. "github.com/shono09835/bosh-cli/v7/release/resource"
	. "github.com/shono09835/bosh-cli/v7/templatescompiler"
)

var _ = Describe("RenderedJob", func() {
	var (
		outBuffer       *bytes.Buffer
		logger          boshlog.Logger
		fs              *fakeboshsys.FakeFileSystem
		releaseJob      boshreljob.Job
		renderedJobPath string
		renderedJob     RenderedJob
	)

	BeforeEach(func() {
		outBuffer = bytes.NewBufferString("")
		logger = boshlog.NewWriterLogger(boshlog.LevelDebug, outBuffer)
		fs = fakeboshsys.NewFakeFileSystem()
		releaseJob = *boshreljob.NewJob(NewResource("fake-job-name", "", nil))
		renderedJobPath = "fake-path"
		renderedJob = NewRenderedJob(releaseJob, renderedJobPath, fs, logger)
	})

	Describe("Job", func() {
		It("returns the release job", func() {
			Expect(renderedJob.Job()).To(Equal(releaseJob))
		})
	})

	Describe("Path", func() {
		It("returns the rendered job path", func() {
			Expect(renderedJob.Path()).To(Equal(renderedJobPath))
		})
	})

	Describe("Delete", func() {
		It("deletes the rendered job path from the file system", func() {
			err := fs.MkdirAll(renderedJobPath, os.ModePerm)
			Expect(err).ToNot(HaveOccurred())

			err = renderedJob.Delete()
			Expect(err).ToNot(HaveOccurred())
			Expect(fs.FileExists(renderedJobPath)).To(BeFalse())
		})

		Context("when deleting from the file system fails", func() {
			JustBeforeEach(func() {
				fs.RemoveAllStub = func(_ string) error {
					return bosherr.Error("fake-delete-error")
				}
			})

			It("returns an error", func() {
				err := renderedJob.Delete()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("fake-delete-error"))
			})
		})
	})

	Describe("DeleteSilently", func() {
		It("deletes the rendered job path from the file system", func() {
			err := fs.MkdirAll(renderedJobPath, os.ModePerm)
			Expect(err).ToNot(HaveOccurred())

			renderedJob.DeleteSilently()
			Expect(fs.FileExists(renderedJobPath)).To(BeFalse())
		})

		Context("when deleting from the file system fails", func() {
			JustBeforeEach(func() {
				fs.RemoveAllStub = func(_ string) error {
					return bosherr.Error("fake-delete-error")
				}
			})

			It("logs the error", func() {
				renderedJob.DeleteSilently()

				log := outBuffer.String()
				Expect(log).To(ContainSubstring("Failed to delete rendered job"))
				Expect(log).To(ContainSubstring("fake-delete-error"))
			})
		})
	})
})
