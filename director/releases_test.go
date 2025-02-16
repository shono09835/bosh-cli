package director_test

import (
	"errors"
	"net/http"
	"os"

	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	semver "github.com/cppforlife/go-semi-semantic/version"
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

	Describe("Releases", func() {
		It("returns releases", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[
  {
    "name": "rel1-name",
    "release_versions": [
      {
        "version": "rel1-ver1",
        "currently_deployed": false,
        "uncommitted_changes": false,
        "commit_hash": "rel1-hash1"
      },
      {
        "version": "rel1-ver2",
        "currently_deployed": true,
        "uncommitted_changes": true,
        "commit_hash": "rel1-hash2"
      }
    ]
  },
  {
    "name": "rel2-name",
    "release_versions": [
      {
        "version": "rel2-ver1",
        "currently_deployed": false,
        "uncommitted_changes": false,
        "commit_hash": "rel2-hash"
      }
    ]
  }
]`),
				),
			)

			rels, err := director.Releases()
			Expect(err).ToNot(HaveOccurred())
			Expect(rels).To(HaveLen(3))

			Expect(rels[0].Name()).To(Equal("rel1-name"))
			Expect(rels[0].Version()).To(Equal(semver.MustNewVersionFromString("rel1-ver1")))
			Expect(rels[0].VersionMark("*")).To(Equal(""))
			Expect(rels[0].CommitHashWithMark("*")).To(Equal("rel1-hash1"))

			Expect(rels[1].Name()).To(Equal("rel1-name"))
			Expect(rels[1].Version()).To(Equal(semver.MustNewVersionFromString("rel1-ver2")))
			Expect(rels[1].VersionMark("*")).To(Equal("*"))
			Expect(rels[1].CommitHashWithMark("*")).To(Equal("rel1-hash2*"))

			Expect(rels[2].Name()).To(Equal("rel2-name"))
			Expect(rels[2].Version()).To(Equal(semver.MustNewVersionFromString("rel2-ver1")))
			Expect(rels[2].VersionMark("*")).To(Equal(""))
			Expect(rels[2].CommitHashWithMark("*")).To(Equal("rel2-hash"))
		})

		It("returns an error for invalid release versions", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"-"}]}]`),
				),
			)

			_, err := director.Releases()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing version for release"))
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/releases"), server)

			_, err := director.Releases()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding releases: Director responded with non-successful status code"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.RespondWith(http.StatusOK, ""),
				),
			)

			_, err := director.Releases()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding releases: Unmarshaling Director response"))
		})
	})

	Describe("FindRelease", func() {
		It("does not return an error", func() {
			_, err := director.FindRelease(NewReleaseSlug("name", "ver"))
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error if version is invalid", func() {
			_, err := director.FindRelease(NewReleaseSlug("name", "-"))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing version for release"))
		})
	})

	Describe("HasRelease", func() {
		var stemcell OSVersionSlug

		act := func() (bool, error) {
			return director.HasRelease("name", "ver", stemcell)
		}

		BeforeEach(func() {
			stemcell = OSVersionSlug{}
		})

		It("returns true if name and version matches and it has source", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases/name", "version=ver"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `{
					  "packages": [{
						"blobstore_id": "123"
					  }]
					}`),
				),
			)

			found, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
		})

		It("returns false if name and version matches but no source for some packages", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases/name", "version=ver"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `{
					  "packages": [{
						"blobstore_id": "ID_EXISTS"
					  },{
						"blobstore_id": ""
					  }]
					}`),
				),
			)

			found, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})

		It("returns false if name and version does not match", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[
  {"name": "name", "release_versions": [{"version": "other-ver"}]},
  {"name": "other-name", "release_versions": [{"version": "ver"}]}
]`),
				),
			)

			found, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})

		It("returns true if name, version and stemcell matches", func() {
			stemcell = NewOSVersionSlug("stemcell-os", "stemcell-ver")

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases/name", "version=ver"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `{
					  "packages": [{
					    "compiled_packages": [{ "stemcell": "stemcell-os/stemcell-ver" }]
					  }]
					}`),
				),
			)

			found, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
		})

		It("returns false for compiled release if name and version match, but director only has source packages", func() {
			stemcell = NewOSVersionSlug("stemcell-os", "stemcell-ver")

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases/name", "version=ver"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `{
					  "packages": [{
					    "compiled_packages": []
					  }]
					}`),
				),
			)

			found, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})

		It("returns false if name, version matches but stemcell does not", func() {
			stemcell = NewOSVersionSlug("stemcell-os", "stemcell-ver")

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases/name", "version=ver"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `{
					  "packages": [{
					    "compiled_packages": [
					    	{ "stemcell": "stemcell-os/stemcell-ver1" },
					    	{ "stemcell": "stemcell-os1/stemcell-ver" }
					    ]
					  }]
					}`),
				),
			)

			found, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})

		It("returns false if name, version matches but not all compiled packages have the matching stemcell", func() {
			stemcell = NewOSVersionSlug("stemcell-os", "stemcell-ver")

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases/name", "version=ver"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `{
					  "packages": [
						{
							"compiled_packages": [
								{ "stemcell": "stemcell-os/stemcell-ver" },
								{ "stemcell": "stemcell-os2/stemcell-ver2" }
							]
						},
						{
							"compiled_packages": [
								{ "stemcell": "stemcell-os2/stemcell-ver2" }
							]
						}
					  ]
					}`),
				),
			)

			found, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/releases"), server)

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding releases: Director responded with non-successful status code"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.RespondWith(http.StatusOK, ""),
				),
			)

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding releases: Unmarshaling Director response"))
		})
	})

	Describe("UploadReleaseURL", func() {
		It("uploads release by URL", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/releases", ""),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"application/json"},
					}),
					ghttp.VerifyBody([]byte(`{"location":"url"}`)),
				),
				"",
				server,
			)

			Expect(director.UploadReleaseURL("url", "", false, false)).ToNot(HaveOccurred())
		})

		It("uploads release by URL with sha1, rebase and fix", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/releases", "rebase=true&fix=true"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"application/json"},
					}),
					ghttp.VerifyBody([]byte(`{"location":"url","sha1":"sha1"}`)),
				),
				"",
				server,
			)

			Expect(director.UploadReleaseURL("url", "sha1", true, true)).ToNot(HaveOccurred())
		})

		It("returns error if URL is empty", func() {
			err := director.UploadReleaseURL("", "", false, false)
			Expect(err).To(Equal(errors.New("Expected non-empty URL")))
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("POST", "/releases"), server)

			err := director.UploadReleaseURL("url", "", false, false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Uploading remote release 'url': Director responded with non-successful status code"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/releases"),
					ghttp.RespondWith(http.StatusOK, ""),
				),
			)

			err := director.UploadReleaseURL("url", "", false, false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Uploading remote release 'url': Unmarshaling Director response"))
		})
	})

	Describe("UploadReleaseFile", func() {
		var (
			file UploadFile
		)

		BeforeEach(func() {
			fs := fakesys.NewFakeFileSystem()
			err := fs.WriteFileString("/file", "content")
			Expect(err).ToNot(HaveOccurred())

			file, err = fs.OpenFile("/file", os.O_RDONLY, 0)
			Expect(err).ToNot(HaveOccurred())
		})

		It("uploads release file", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/releases", ""),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type":   []string{"application/x-compressed"},
						"Content-Length": []string{"7"},
					}),
					ghttp.VerifyBody([]byte("content")),
				),
				"",
				server,
			)

			Expect(director.UploadReleaseFile(file, false, false)).ToNot(HaveOccurred())
		})

		It("uploads release file with rebase and fix", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/releases", "rebase=true&fix=true"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type":   []string{"application/x-compressed"},
						"Content-Length": []string{"7"},
					}),
					ghttp.VerifyBody([]byte("content")),
				),
				"",
				server,
			)

			Expect(director.UploadReleaseFile(file, true, true)).ToNot(HaveOccurred())
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("POST", "/releases"), server)

			err := director.UploadReleaseFile(file, false, false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Uploading release file: Director responded with non-successful status code"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/releases"),
					ghttp.RespondWith(http.StatusOK, ""),
				),
			)

			err := director.UploadReleaseFile(file, false, false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Uploading release file: Unmarshaling Director response"))
		})
	})
})

var _ = Describe("Release", func() {
	var (
		director Director
		release  Release
		server   *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()

		var err error

		release, err = director.FindRelease(NewReleaseSlug("name", "ver"))
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Name", func() {
		It("returns name", func() {
			Expect(release.Name()).To(Equal("name"))
		})
	})

	Describe("Exists", func() {
		It("returns true if release exists", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
			)

			exists, err := release.Exists()
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})

		It("returns false if release does not exist", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"other-name","release_versions":[{"version":"other-ver"}]}]`),
				),
			)

			exists, err := release.Exists()
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})

		It("returns false and an error when failing to get the release slug", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/releases"), server)

			exists, err := release.Exists()
			Expect(err).To(HaveOccurred())
			Expect(exists).To(BeFalse())
		})
	})

	Describe("Delete", func() {
		It("succeeds deleting", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", "/releases/name", "version=ver"),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				"",
				server,
			)

			Expect(release.Delete(false)).ToNot(HaveOccurred())
		})

		It("succeeds deleting with force flag", func() {
			ConfigureTaskResult(ghttp.VerifyRequest("DELETE", "/releases/name", "version=ver&force=true"), "", server)

			Expect(release.Delete(true)).ToNot(HaveOccurred())
		})

		It("succeeds even if error occurrs if release no longer exist", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/releases/name"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, "[]"),
				),
			)

			Expect(release.Delete(false)).ToNot(HaveOccurred())
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

			err := release.Delete(false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Deleting release or series 'name[/ver]': Director responded with non-successful status code"))
		})

		It("returns delete error if response is non-200 and release still exists", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/releases/name"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/releases"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name":"name","release_versions":[{"version":"ver"}]}]`),
				),
			)

			err := release.Delete(false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Deleting release or series 'name[/ver]': Director responded with non-successful status code"))
		})
	})
})
