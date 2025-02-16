package release_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/cpi/release"
	boshrel "github.com/shono09835/bosh-cli/v7/release"
	boshjob "github.com/shono09835/bosh-cli/v7/release/job"
	fakerel "github.com/shono09835/bosh-cli/v7/release/releasefakes"
	. "github.com/shono09835/bosh-cli/v7/release/resource"
)

var _ = Describe("Validator", func() {
	var cpiReleaseJobName = "fake-cpi-release-job-name"

	It("validates a valid release without error", func() {
		job := boshjob.NewJob(NewResourceWithBuiltArchive(
			"fake-cpi-release-job-name", "fake-job-1-fingerprint", "", "fake-job-1-sha"))

		job.Templates = map[string]string{"cpi.erb": "bin/cpi"}

		release := &fakerel.FakeRelease{
			NameStub:    func() string { return "fake-release-name" },
			VersionStub: func() string { return "fake-release-version" },

			FindJobByNameStub: func(name string) (boshjob.Job, bool) {
				Expect(name).To(Equal(job.Name()))
				return *job, true
			},
		}

		validator := NewValidator()

		err := validator.Validate(release, cpiReleaseJobName)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when the cpi job is not present", func() {
		var validator Validator
		var release *fakerel.FakeRelease

		BeforeEach(func() {
			job := boshjob.NewJob(NewResourceWithBuiltArchive(
				"non-cpi-job", "fake-job-1-fingerprint", "", "fake-job-1-sha"))

			job.Templates = map[string]string{"cpi.erb": "bin/cpi"}

			release = &fakerel.FakeRelease{
				NameStub:    func() string { return "fake-release-name" },
				VersionStub: func() string { return "fake-release-version" },

				FindJobByNameStub: func(name string) (boshjob.Job, bool) {
					Expect(name).To(Equal(cpiReleaseJobName))
					return boshjob.Job{}, false
				},
			}

			validator = NewValidator()
		})

		It("returns an error that the cpi job is not present", func() {
			err := validator.Validate(release, cpiReleaseJobName)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"CPI release must contain specified job 'fake-cpi-release-job-name'"))
		})
	})

	Context("when the templates are missing a bin/cpi target", func() {
		var validator Validator
		var release boshrel.Release

		BeforeEach(func() {
			job := boshjob.NewJob(NewResourceWithBuiltArchive(
				"fake-cpi-release-job-name", "fake-job-1-fingerprint", "", "fake-job-1-sha"))

			job.Templates = map[string]string{"cpi.erb": "nonsense"}

			release = &fakerel.FakeRelease{
				NameStub:    func() string { return "fake-release-name" },
				VersionStub: func() string { return "fake-release-version" },

				FindJobByNameStub: func(name string) (boshjob.Job, bool) {
					Expect(name).To(Equal(job.Name()))
					return *job, true
				},
			}

			validator = NewValidator()
		})

		It("returns an error that the bin/cpi template target is missing", func() {
			err := validator.Validate(release, cpiReleaseJobName)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Specified CPI release job 'fake-cpi-release-job-name' must contain a template that renders to target 'bin/cpi'"))
		})
	})
})
