package director_test

import (
	"net/http"

	. "github.com/shono09835/bosh-cli/v7/director"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Director", func() {
	var (
		director   Director
		deployment Deployment
		server     *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()

		var err error

		deployment, err = director.FindDeployment("dep")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("DeleteVM", func() {
		It("deletes VM", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", "/vms/cid"),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				"",
				server,
			)
			err := deployment.DeleteVM("cid")
			Expect(err).ToNot(HaveOccurred())
		})

		It("does url encoding for cid", func() {
			var verifyRawPath = func(path string) http.HandlerFunc {
				return func(w http.ResponseWriter, req *http.Request) {
					Expect(req.RequestURI).To(Equal(path))
				}
			}

			ConfigureTaskResult(
				ghttp.CombineHandlers(
					verifyRawPath("/vms/cid%3Bcid"),
					ghttp.VerifyRequest("DELETE", "/vms/cid;cid"),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				"",
				server,
			)
			err := deployment.DeleteVM("cid;cid")
			Expect(err).ToNot(HaveOccurred())
		})

		It("succeeds even if error occurrs if VM no longer exists", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/vms/cid"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/vms"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[]`),
				),
			)

			err := deployment.DeleteVM("cid")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns delete error if listing VMs fails", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/vms/cid"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/vms"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			err := deployment.DeleteVM("cid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Deleting VM 'cid'"))
		})

		It("returns delete error if response is non-200 and VM still exists", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/vms/cid"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/vms"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"vm_cid": "cid"}]`),
				),
			)

			err := deployment.DeleteVM("cid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Deleting VM 'cid'"))
		})
	})
})
