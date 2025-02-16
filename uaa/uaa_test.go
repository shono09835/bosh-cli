package uaa_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/shono09835/bosh-cli/v7/uaa"
)

var _ = Describe("UAA", func() {
	var (
		uaa    UAA
		server *ghttp.Server
	)

	BeforeEach(func() {
		uaa, server = BuildServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("RefreshTokenGrant", func() {
		It("returns a new access token that can only be refreshed", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.VerifyBody([]byte("grant_type=refresh_token&refresh_token=refresh-token")),
					ghttp.VerifyHeader(http.Header{"Content-Type": []string{"application/x-www-form-urlencoded"}}),
					ghttp.RespondWith(http.StatusOK, `{
                 		"token_type": "new-bearer",
                 		"access_token": "new-access-token",
										"refresh_token": "new-refresh-token"
	                }`),
				),
			)

			newToken, err := uaa.RefreshTokenGrant("refresh-token")
			Expect(err).ToNot(HaveOccurred())
			Expect(newToken.Type()).To(Equal("new-bearer"))
			Expect(newToken.Value()).To(Equal("new-access-token"))

			newRefreshToken, refreshable := newToken.(RefreshableAccessToken)
			Expect(refreshable).To(BeTrue())
			Expect(newRefreshToken.RefreshValue()).To(Equal("new-refresh-token"))
		})

		It("returns error if token response is non-200", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWith(http.StatusBadRequest, ``),
				),
			)

			_, err := uaa.RefreshTokenGrant("refresh-token")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("UAA responded with non-successful status code"))
		})

		It("returns error if token cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			_, err := uaa.RefreshTokenGrant("refresh-token")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unmarshaling UAA response"))
		})
	})

	Describe("ClientCredentialsGrant", func() {
		It("obtains client token", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.VerifyBody([]byte("grant_type=client_credentials")),
					ghttp.VerifyHeader(http.Header{"Content-Type": []string{"application/x-www-form-urlencoded"}}),
					ghttp.VerifyBasicAuth("client", "client-secret"),
					ghttp.VerifyHeader(http.Header{
						"Accept": []string{"application/json"},
					}),
					ghttp.RespondWith(http.StatusOK, `{
                 		"token_type": "bearer",
                 		"access_token": "access-token"
	                }`),
				),
			)

			token, err := uaa.ClientCredentialsGrant()
			Expect(err).ToNot(HaveOccurred())
			Expect(token.Type()).To(Equal("bearer"))
			Expect(token.Value()).To(Equal("access-token"))
		})

		It("returns error if prompts response in non-200", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWith(http.StatusBadRequest, ``),
				),
			)

			_, err := uaa.ClientCredentialsGrant()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("UAA responded with non-successful status code"))
		})

		It("returns error if prompts cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			_, err := uaa.ClientCredentialsGrant()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unmarshaling UAA response"))
		})
	})

	Describe("OwnerPasswordCredentialsGrant", func() {
		It("obtains access token based on prompt answers", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.VerifyBody([]byte("grant_type=password&key1=ans1&key2=ans2")),
					ghttp.VerifyHeader(http.Header{"Content-Type": []string{"application/x-www-form-urlencoded"}}),
					ghttp.VerifyBasicAuth("client", "client-secret"),
					ghttp.VerifyHeader(http.Header{
						"Accept": []string{"application/json"},
					}),
					ghttp.RespondWith(http.StatusOK, `{
                 		"token_type": "bearer",
                 		"access_token": "access-token",
                 		"refresh_token": "refresh-token"
	                }`),
				),
			)

			answers := []PromptAnswer{
				{Key: "key1", Value: "ans1"},
				{Key: "key2", Value: "ans2"},
			}

			token, err := uaa.OwnerPasswordCredentialsGrant(answers)
			Expect(err).ToNot(HaveOccurred())
			Expect(token.Type()).To(Equal("bearer"))
			Expect(token.Value()).To(Equal("access-token"))

			newRefreshToken, refreshable := token.(RefreshableAccessToken)
			Expect(refreshable).To(BeTrue())
			Expect(newRefreshToken.RefreshValue()).To(Equal("refresh-token"))
		})

		It("returns error if prompts response in non-200", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWith(http.StatusBadRequest, ``),
				),
			)

			_, err := uaa.OwnerPasswordCredentialsGrant(nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("UAA responded with non-successful status code"))
		})

		It("returns error if prompts cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			_, err := uaa.OwnerPasswordCredentialsGrant(nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unmarshaling UAA response"))
		})
	})
})
