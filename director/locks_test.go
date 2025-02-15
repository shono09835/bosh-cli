package director_test

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/shono09835/bosh-cli/v7/director"
)

var _ = Describe("Director", func() {
	var (
		director Director
		server   *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Locks", func() {
		It("returns locks", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/locks"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[
	{
		"type": "deployment",
		"resource": ["some-deployment-name"],
		"timeout": "1443889622.9964118",
		"task_id": "1234567"
	},
	{
		"type": "release",
		"resource": ["some-release-name", "123"],
		"timeout": "1443889622.9964118",
		"task_id": "123456"
	}
]`),
				),
			)

			locks, err := director.Locks()
			Expect(err).ToNot(HaveOccurred())
			Expect(locks).To(Equal([]Lock{
				{
					Type:      "deployment",
					Resource:  []string{"some-deployment-name"},
					ExpiresAt: time.Unix(1443889622, 0).UTC(),
					TaskID:    "1234567",
				},
				{
					Type:      "release",
					Resource:  []string{"some-release-name", "123"},
					ExpiresAt: time.Unix(1443889622, 0).UTC(),
					TaskID:    "123456",
				},
			}))
		})

		It("returns error if lock timeout cannot be parsed", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/locks"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[
	{
		"type": "deployment",
		"resource": ["some-deployment-name"],
		"timeout": "invalid",
		"task_id": "123456"
	}
]`),
				),
			)

			_, err := director.Locks()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Converting timeout 'invalid' to float"))
		})

		It("returns error if info response in non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/locks"), server)

			_, err := director.Locks()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding locks: Director responded with non-successful status code"))
		})

		It("returns error if info cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/locks"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			_, err := director.Locks()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding locks: Unmarshaling Director response"))
		})
	})
})
