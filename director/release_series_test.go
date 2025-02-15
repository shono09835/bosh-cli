package director_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/shono09835/bosh-cli/v7/director"
)

var _ = Describe("Director", func() {
	Describe("FindReleaseSeries", func() {
		It("does not return an error", func() {
			director, server := BuildServer()
			defer server.Close()

			_, err := director.FindReleaseSeries(NewReleaseSeriesSlug("name"))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("ReleaseSeries", func() {
	var (
		director Director
		series   ReleaseSeries
		server   *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()

		var err error

		series, err = director.FindReleaseSeries(NewReleaseSeriesSlug("name"))
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Name", func() {
		It("returns name", func() {
			Expect(series.Name()).To(Equal("name"))
		})
	})

	Describe("Exists", func() {
		It("returns true when release series exists", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name": "name"}]`),
				),
			)
			exist, err := series.Exists()

			Expect(err).ToNot(HaveOccurred())
			Expect(exist).To(Equal(true))
		})

		It("returns false when release series does not exist", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name": "other-name"}]`),
				),
			)
			exist, err := series.Exists()

			Expect(err).ToNot(HaveOccurred())
			Expect(exist).To(Equal(false))
		})

		It("returns false and an error when failing to get the release series", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/releases"), server)
			exist, err := series.Exists()

			Expect(err).To(HaveOccurred())
			Expect(exist).To(Equal(false))
		})
	})

	Describe("Delete", func() {
		It("succeeds deleting", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", "/releases/name", ""),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				"",
				server,
			)

			Expect(series.Delete(false)).ToNot(HaveOccurred())
		})

		It("succeeds deleting with force flag", func() {
			ConfigureTaskResult(ghttp.VerifyRequest("DELETE", "/releases/name", "force=true"), "", server)

			Expect(series.Delete(true)).ToNot(HaveOccurred())
		})

		It("succeeds even if error occurrs if release series no longer exist", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/releases/name"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, "[]"),
				),
			)

			Expect(series.Delete(false)).ToNot(HaveOccurred())
		})

		It("returns delete error if listing releases fails", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/releases/name"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, ""),
				),
			)

			err := series.Delete(false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Deleting release or series 'name[/]': Director responded with non-successful status code"))
		})

		It("returns delete error if response is non-200 and release still exists", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/releases/name"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name": "name"}]`),
				),
			)

			err := series.Delete(false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Deleting release or series 'name[/]': Director responded with non-successful status code"))
		})
	})
})
