package releasedir_test

import (
	"errors"
	"path/filepath"
	"strings"

	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	fakeuuid "github.com/cloudfoundry/bosh-utils/uuid/fakes"
	semver "github.com/cppforlife/go-semi-semantic/version"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	boshman "github.com/shono09835/bosh-cli/v7/release/manifest"
	fakerel "github.com/shono09835/bosh-cli/v7/release/releasefakes"
	. "github.com/shono09835/bosh-cli/v7/releasedir"
	fakereldir "github.com/shono09835/bosh-cli/v7/releasedir/releasedirfakes"
)

var _ = Describe("FSReleaseIndex", func() {
	var (
		reporter *fakereldir.FakeReleaseIndexReporter
		uuidGen  *fakeuuid.FakeGenerator
		fs       *fakesys.FakeFileSystem
		index    FSReleaseIndex
	)

	BeforeEach(func() {
		reporter = &fakereldir.FakeReleaseIndexReporter{}
		uuidGen = &fakeuuid.FakeGenerator{}
		fs = fakesys.NewFakeFileSystem()
		index = NewFSReleaseIndex("index-name", filepath.Join("/", "dir"), reporter, uuidGen, fs)
	})

	Describe("LastVersion", func() {
		It("returns nil when there is no index file", func() {
			ver, err := index.LastVersion("name")
			Expect(err).ToNot(HaveOccurred())
			Expect(ver).To(BeNil())
		})

		It("returns nil when index file is empty", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "")
			Expect(err).ToNot(HaveOccurred())

			ver, err := index.LastVersion("name")
			Expect(err).ToNot(HaveOccurred())
			Expect(ver).To(BeNil())
		})

		It("returns greater version", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), `---
builds:
  uuid1: {version: "1.1"}
  uuid2: {version: "1"}
format-version: "2"`)
			Expect(err).ToNot(HaveOccurred())

			ver, err := index.LastVersion("name")
			Expect(err).ToNot(HaveOccurred())
			Expect(ver.String()).To(Equal(semver.MustNewVersionFromString("1.1").String()))
		})

		It("returns error if version cannot be parsed", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), `---
builds:
  uuid2: {version: "-"}
format-version: "2"`)
			Expect(err).ToNot(HaveOccurred())

			_, err = index.LastVersion("name")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing release versions"))
		})

		It("returns error if name is empty", func() {
			_, err := index.LastVersion("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected non-empty release name"))
		})

		It("returns error if index file cannot be unmarshalled", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "-")
			Expect(err).ToNot(HaveOccurred())

			_, err = index.LastVersion("name")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("line 1"))
		})

		It("returns error if reading index file fails", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "")
			Expect(err).ToNot(HaveOccurred())
			fs.ReadFileError = errors.New("fake-err")

			_, err = index.LastVersion("name")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})

	Describe("Contains", func() {
		var (
			release *fakerel.FakeRelease
		)

		BeforeEach(func() {
			release = &fakerel.FakeRelease{}
			release.NameReturns("name")
			release.VersionReturns("ver1")
		})

		It("returns false when there is no index file", func() {
			exists, err := index.Contains(release)
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})

		It("returns false when index file is empty", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "")
			Expect(err).ToNot(HaveOccurred())

			exists, err := index.Contains(release)
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})

		It("returns true if version is exists", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), `---
builds:
  uuid1: {version: "1.1"}
  uuid2: {version: "ver1"}
format-version: "2"`)
			Expect(err).ToNot(HaveOccurred())

			exists, err := index.Contains(release)
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})

		It("returns false if version is not exists", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), `---
builds:
  uuid1: {version: "1.1"}
format-version: "2"`)
			Expect(err).ToNot(HaveOccurred())

			exists, err := index.Contains(release)
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})

		It("returns error if name is empty", func() {
			release.NameReturns("")

			_, err := index.Contains(release)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected non-empty release name"))
		})

		It("returns error if version is empty", func() {
			release.VersionReturns("")

			_, err := index.Contains(release)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected non-empty release version"))
		})

		It("returns error if index file cannot be unmarshalled", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "-")
			Expect(err).ToNot(HaveOccurred())

			_, err = index.Contains(release)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("line 1"))
		})

		It("returns error if reading index file fails", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "")
			Expect(err).ToNot(HaveOccurred())
			fs.ReadFileError = errors.New("fake-err")

			_, err = index.Contains(release)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})

	Describe("Add", func() {
		var (
			manifest boshman.Manifest
		)

		BeforeEach(func() {
			manifest.Name = "name"
			manifest.Version = "ver1"
			uuidGen.GeneratedUUID = "new-uuid"
		})

		It("saves manifest and adds version entry when there is no index file", func() {
			err := index.Add(manifest)
			Expect(err).ToNot(HaveOccurred())

			Expect(fs.ReadFileString(filepath.Join("/", "dir", "name", "name-ver1.yml"))).To(Equal(`name: name
version: ver1
commit_hash: ""
uncommitted_changes: false
`))

			Expect(fs.ReadFileString(filepath.Join("/", "dir", "name", "index.yml"))).To(Equal(`builds:
  new-uuid:
    version: ver1
format-version: "2"
`))

			name, desc, err := reporter.ReleaseIndexAddedArgsForCall(0)
			Expect(name).To(Equal("index-name"))
			Expect(desc).To(Equal("name/ver1"))
			Expect(err).To(BeNil())
		})

		It("saves manifest and adds version entry", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), `---
builds:
  uuid: {version: "1.1"}
format-version: "2"
`)
			Expect(err).ToNot(HaveOccurred())

			err = index.Add(manifest)
			Expect(err).ToNot(HaveOccurred())

			Expect(fs.ReadFileString(filepath.Join("/", "dir", "name", "name-ver1.yml"))).To(Equal(`name: name
version: ver1
commit_hash: ""
uncommitted_changes: false
`))

			Expect(fs.ReadFileString(filepath.Join("/", "dir", "name", "index.yml"))).To(Equal(`builds:
  new-uuid:
    version: ver1
  uuid:
    version: "1.1"
format-version: "2"
`))

			name, desc, err := reporter.ReleaseIndexAddedArgsForCall(0)
			Expect(name).To(Equal("index-name"))
			Expect(desc).To(Equal("name/ver1"))
			Expect(err).To(BeNil())
		})

		It("returns and reports error if writing manifest fails", func() {
			fs.WriteFileErrors[filepath.Join("/", "dir", "name", "name-ver1.yml")] = errors.New("fake-err")

			err := index.Add(manifest)
			Expect(err).To(HaveOccurred())

			name, desc, err := reporter.ReleaseIndexAddedArgsForCall(0)
			Expect(name).To(Equal("index-name"))
			Expect(desc).To(Equal("name/ver1"))
			Expect(err).ToNot(BeNil())
		})

		It("returns and reports error if writing index fails", func() {
			fs.WriteFileErrors[filepath.Join("/", "dir", "name", "index.yml")] = errors.New("fake-err")

			err := index.Add(manifest)
			Expect(err).To(HaveOccurred())

			name, desc, err := reporter.ReleaseIndexAddedArgsForCall(0)
			Expect(name).To(Equal("index-name"))
			Expect(desc).To(Equal("name/ver1"))
			Expect(err).ToNot(BeNil())
		})

		It("returns error if generating uuid fails", func() {
			uuidGen.GenerateError = errors.New("fake-err")

			err := index.Add(manifest)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("returns error if version is exists", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), `---
builds:
  uuid1: {version: "1.1"}
  uuid2: {version: "ver1"}
format-version: "2"`)
			Expect(err).ToNot(HaveOccurred())

			err = index.Add(manifest)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Release version 'ver1' already exists"))
		})

		It("returns error if name is empty", func() {
			manifest.Name = ""

			err := index.Add(manifest)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected non-empty release name"))
		})

		It("returns error if version is empty", func() {
			manifest.Version = ""

			err := index.Add(manifest)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected non-empty release version"))
		})

		It("returns error if index file cannot be unmarshalled", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "-")
			Expect(err).ToNot(HaveOccurred())

			err = index.Add(manifest)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("line 1"))
		})

		It("returns error if reading index file fails", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), "")
			Expect(err).ToNot(HaveOccurred())
			fs.ReadFileError = errors.New("fake-err")

			err = index.Add(manifest)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("does not reorder keys needlessly", func() {
			err := fs.WriteFileString(filepath.Join("/", "dir", "name", "index.yml"), fsReleaseIndexSortingFixture)
			Expect(err).ToNot(HaveOccurred())

			uuidGen.GeneratedUUID = "new-uuid"
			manifest.Version = "new"

			err = index.Add(manifest)
			Expect(err).ToNot(HaveOccurred())

			afterFirstSort, err := fs.ReadFileString(filepath.Join("/", "dir", "name", "index.yml"))
			Expect(err).ToNot(HaveOccurred())

			Expect(afterFirstSort).ToNot(Equal(fsReleaseIndexSortingFixture)) // sanity check

			uuidGen.GeneratedUUID = "another-uuid"
			manifest.Version = "another"

			err = index.Add(manifest)
			Expect(err).ToNot(HaveOccurred())

			after, err := fs.ReadFileString(filepath.Join("/", "dir", "name", "index.yml"))
			Expect(err).ToNot(HaveOccurred())

			Expect(after).ToNot(Equal(afterFirstSort)) // sanity check
			Expect(strings.Replace(after, "  another-uuid:\n    version: another\n", "", 1)).To(Equal(afterFirstSort))
		})
	})

	Describe("ManifestPath", func() {
		It("returns path to a manifest", func() {
			Expect(index.ManifestPath("name", "ver1")).To(Equal(filepath.Join("/", "dir", "name", "name-ver1.yml")))
		})
	})
})

// Fixture needs to be long because natural sort may succeed for smaller sizes
const fsReleaseIndexSortingFixture = `
builds:
  0a09e695-c0d6-4283-ad1e-132a336095d6:
    version: "184"
  0be3304e-c217-4b0b-8f21-2b4ae1811f02:
    version: "96"
  0c3ce4ae-2ad7-4b70-bf68-3a11c199694d:
    version: "140"
  0c923cc3-49ee-4376-9033-c51f7380794b:
    version: "256.7"
  0d469717-3f49-4e68-8472-7e11ddcd8a44:
    version: "132"
  0dcbed5b-7245-4223-b99a-35d7b5ddade4:
    version: "189"
  0e3ed48f-d647-44c3-804b-e44bda1f9e55:
    version: "105"
  1abbb9d95dd0121a238e319580cda00dcc9d6e82:
    version: "86"
  1bb45287-c018-4e39-86a1-6b5a65c49668:
    version: "97"
  1bc91aed-94f9-40db-962a-2dae64bcbd10:
    version: "130"
  1eb2cc5facfa83077237c30d3ef8a875ba76dda6:
    version: "25"
  2cd44cf882745de94eb4b470d9f5659f6d752388:
    version: "2"
  2e47bdf8-79fb-4b5f-8a18-8c00a08df116:
    version: "257"
  2e2889bbd5168b5c8daff2a42ae2bf7aaec22126:
    version: "10"
  2efbb96d-7e98-42b0-a0aa-1e20f090a755:
    version: "183"
  2f343a15-b377-4761-a72c-43f06ceed112:
    version: "243"
  002f99bb-3087-461a-bd93-9cef673f04a7:
    version: "237"
  3b1fd3e4-7756-459d-8930-68c8fcbd6a3f:
    version: "247"
  3b6704a1-f33e-49eb-924a-af53f9bd47e4:
    version: "104"
  3bb433e7-370d-444a-8a27-14871de4e7c0:
    version: "145"
  3cc94519942dd522508a487a3282956d84c3b21c:
    version: "8"
  3d933826-e485-47d4-938a-0da83afba887:
    version: "211"
  3f36e5523e474c55424543334971ddbf578a354f:
    version: "24"
  4a1784f3-90d3-4e10-80a9-7084dfeae145:
    version: "129"
  4a2233a4-edf6-4095-a91b-bbd5386b240e:
    version: "255.9"
  4b3c83cbb984ef05560e51077e42aa91567eb91b:
    version: "47"
  4b5ddee8-a565-47c0-a962-57544778769c:
    version: "203"
  4b375c21-0fe1-43e7-b588-c77e8fd07d72:
    version: "151"
  4b685ba7-d4b2-4d4a-84fd-a728dce65bd5:
    version: "239"
  4c43ec6fe8ac88358c106675b0979b0b32b60db0:
    version: "19"
  4d28e0be-b536-4d3a-a5c8-6003cffa0063:
    version: "125"
  4d48c371-a4fb-4887-b00c-1c2496fc530c:
    version: "255.8"
  4df46a47ec8514b03157cb5bfab95c79b2e10965:
    version: "16"
  4e2a2400-1659-4170-b1a8-926638d6ed73:
    version: "114"
  4e2d1d17-3965-4022-8323-c016036866bf:
    version: "167"
  4e05dc91f07ea0fbaef82c09a0f67f4017c1a6db:
    version: "34"
  4ef38824-c4e8-46fc-8c09-65d6352ee018:
    version: "255.3"
  4f1c7a7af09b3564c53f1ce978832c99ee284b48:
    version: "57"
  04ed089c-86ec-4b3f-b77f-2dfaadc30101:
    version: "116"
  51e5280d-c00c-4fbe-9854-1650e21c23e0:
    version: "255"
  565db893e00b6743e688026817a9f318bb434035:
    version: "37"
  56b8711d637ca966d83dc70785f382db04b0d2ed:
    version: "74"
  56d23a64-5615-49ae-9142-66a49191bd11:
    version: "138"
  515e1ed5-c6db-4548-8e45-a2fd857dd96e:
    version: "218"
  5a691f4cca5237a2dcc09eb3b3fd95fe65f2c434:
    version: "90"
  5b2fac9f-7e26-4e03-a71a-083248a6d923:
    version: "200"
  5b8cc42918a1ba56d8d08cb7efaa4ab88312cf12:
    version: "75"
  5c0a4550-88ce-450e-a025-6664a4f3ffaa:
    version: "99"
  5c7e6536ca86a51a5cb2982d5c7a6bf735a28587:
    version: "6"
  5e35aa8c-b86b-4f12-8008-f55b448c996e:
    version: "245"
  5e548ed3-311c-4058-ac72-82aa8c6e77f1:
    version: "255.1"
  5ebdc4fb89b80130073a1a9e4e0f0f9d32d6b90a:
    version: "93"
  5f75bc56-4974-43a0-9c30-761f09ccef4b:
    version: "133"
  61c073cd-1d75-4278-6332-fef80546f494:
    version: "257.23"
  8a3f7bc77714e234002ab5f6b4e1b89a79db2060:
    version: "83"
  8b15bb32-d191-44b3-8914-883a195b3805:
    version: "223"
  8df8e9ef-fbc6-42fc-b689-3d82d4f8532d:
    version: "154"
  8eb56d4e-8fd5-48d3-b0fc-4801840d83f1:
    version: "256.10"
  8f9964a522c45c13c1c5c678a03c3314ae016a9a:
    version: "18"
  9b6ba124-d4ba-4749-abbf-be58aff92034:
    version: "136"
  9b12286c2f3b83e091190ccabc83e08d91371b69:
    version: "78"
  9b961347-9053-4fd1-8906-dce81072b44f:
    version: "127"
  9bfc22ad-642c-4271-900a-d070556ea2dc:
    version: "157"
  9dfd44e4-58d1-473d-91d8-95fd7af980dd:
    version: "139"
  9e2cfee4-23d2-4333-826f-e3643397925d:
    version: "149"
  9e944111-3908-4def-a928-8a102515881b:
    version: "159"
  09b2f0f1123984e6de289a5276c4f327ff6a4431:
    version: "79"
  10c1abeb97555968f827625b88b82d37fa6d5631:
    version: "70"
  11d6cbdbe2b2f7a7d89c3b238a73753b5aef56cf:
    version: "54"
  15b5afb5-ab70-42a5-bb26-62cd41b0ff83:
    version: "215"
  15d07e9c-e4ce-4980-936e-30f21ffb3e0f:
    version: "214"
  16fdbd44-0cd7-4b3d-939d-a47166146a2a:
    version: "124"
  19e85380-7867-4ad3-a222-fc973f414762:
    version: "255.5"
  27d3af15-dca6-4dd6-9008-83476d87fd4d:
    version: "257.1"
  33e8dd59-6755-4cd7-a3c9-524e5ebe9f71:
    version: "232"
  37cae4d6-12a6-424f-bc6e-d0c8660845a6:
    version: "226"
  40badd7e05dc716feaef7fa7b643ab23b12bc40d:
    version: "73"
  45d5d809-366d-400b-a5c8-92c13be4ecf3:
    version: "217"
  45f55b5020e1002a0eec171f92a4f3543e192e61:
    version: "81"
  47bf7323-8c26-4b70-883d-2a9cc98d707e:
    version: "196"
  63fc1567-7c24-4191-8593-09bc492ef628:
    version: "175"
  70ea18ac-b5f8-4995-8ea9-b33e68b8d5e2:
    version: "173"
  75fc1bea-06a9-4c86-a45f-c04eb960136e:
    version: "187"
  7ab647f56af79aa55544514ace3b100374d9fa0c:
    version: "11"
  7ad3ae21c4ed8b13c89231bafd7488346a2d33cd:
    version: "87"
  7bfc2ae8-f9bd-4801-b116-c37261b31537:
    version: "255.11"
  7cccb6b9-e265-4ea3-8136-e87f09c06dbc:
    version: "238"
  7ccf7897e726ffcd0609a4355034e7500e0b9992:
    version: "5"
  7d958a6c-7c5e-4879-9099-18e8b1f56ce2:
    version: "241"
  80cebae99ee55ecfd1ba543612ba440400b66d6a:
    version: "64"
  84bf30ee-5717-4f43-9c35-45c0b02c963a:
    version: "256"
  84fcfb0e-6a95-4ca4-b501-de607be32277:
    version: "171"
  90d22154-7639-40e5-abdb-9b179056f3e7:
    version: "194"
  92de234b-2361-4fa6-b8e5-48db2c77efa7:
    version: "121"
  93c66707a644f0704cf95ca67235d30ec4ac1706:
    version: "61"
  106d1c7461492a62f6d3ac9e069906183f9839f2:
    version: "20"
  140c3d64ece43e7b182c3d5704553c52b39e08d6:
    version: "80"
  249c9868-7a8b-4a75-9341-4ed3c3f91c81:
    version: "147"
  274cbf9640a23e5287c692bc44f7e155545e199e:
    version: "26"
  328bbda7-5539-47bd-89bf-fac5f004f0dc:
    version: "235"
  335f52e9-4f9d-43fb-bca6-de7bf1cc41cc:
    version: "208"
  424cb5ab917266d52cacd0f5d48a7961bfb18f89:
    version: "59"
  425b86ac-b97e-4905-82ce-1d60597c6bf8:
    version: "152"
  442ab1a9-5598-4962-a8ae-8f9fc027d79f:
    version: "250"
  462af3279520d4041ab444c21bc442865a155f3b:
    version: "9"
  475c9a93753be28b09398dd529e3bb143931a00b:
    version: "36"
  478fc29c04c8e510a6ca0c7e89573b04084cbf89:
    version: "38"
  483be75c-bc09-465d-bec2-3fc2f3a1d564:
    version: "257.14"
  628dc87805e060af7506bc5f5df4ababa8873eaf:
    version: "22"
  00650b34-7377-444e-8317-5033f892602e:
    version: "255.7"
  669d530b-a551-475d-9267-241d16bb2641:
    version: "202"
  66da66653bd20753c15ea5a04c26b857c36eeb98:
    version: "46"
  6b93e0d4-f067-4328-a2c1-9ff3cb2bb808:
    version: "166"
  6c1f79fae1eb0d9b507bfbe7c7651b66b46ab6a5:
    version: "62"
  6d93d456a2d1b7d698893bf826cbd43907740f93:
    version: "68"
  6e2c768f-a6b3-4ea5-a4a6-55ecb1eb4141:
    version: "163"
  6e39c67b-82f4-4b09-84cc-2c8b71889de8:
    version: "248"
  6eb8b63373d273160b71470b1621ccd347a1c429:
    version: "50"
  6ee52fcf-d3a0-4ba5-9dd2-5ca8738875d7:
    version: "225"
  6f03636b-8735-4a90-8c65-e57b88e498d1:
    version: "190"
  6f0cc9b8-be9f-4e91-a354-35a9b69433f7:
    version: "160"
  6fb0e10df7413b89ad6b34e7e92f2fc5f8c50c51:
    version: "13"
  818f7e1675ec75cf1280c3acc76306b28a32708f:
    version: "84"
  824d8f31-bc7a-4a1f-aac2-aaf612d0b016:
    version: "186"
  839b8c7a-275e-46eb-95d4-1db6d69be656:
    version: "221"
  922df222-bbd8-4dec-9cde-efca62eb1d40:
    version: "210"
  971c567baf65344247252f7149887826dcf48521:
    version: "91"
  1386de6b29c9e67be05850f8fa56fccdd98d1783:
    version: "60"
  2011ffa8-cd0d-460b-b1ca-606e5eeebd1d:
    version: "207"
  2946c7ac-ee50-437d-a21a-9e969cfb63bd:
    version: "229"
  3656d798-fa97-4ecc-8f85-05b602ec22d6:
    version: "107"
  4271d8f3-ce32-4a50-b0e2-a1bf0be4fdd8:
    version: "195"
  5347b0a8-7569-42ca-a1ae-02317e584600:
    version: "246"
  5801ac09-3c64-46ba-a39e-3b55336e0036:
    version: "180"
  6295f56b-1d46-4ddb-9491-4d1369f7c77c:
    version: "148"
  6982c119-c61c-441a-9e2d-a91fb22a1c5f:
    version: "141"
  7231a284-2989-4767-9501-618978362f5b:
    version: "227"
  8044cc21-21fd-49eb-903f-cc625b21c4bf:
    version: "257.3"
  8263f0f077b494e90efe16b7fe311b5001a91c9c:
    version: "30"
  9105f65b-9977-402a-bcf1-2341c30de953:
    version: "101"
  9150c4bd-ca2b-4fb3-9997-dc4124a63a94:
    version: "228"
  9296c8d5-2188-454f-8db0-2c754c831d22:
    version: "179"
  9565cc7d-5f58-4a0a-ac67-77094433ceeb:
    version: "108"
  9792e9c92a4a3e9af8f713c49aab48dc2b3698de:
    version: "58"
  9936c4bc-aee0-48dd-8cbf-9cc6d8d1953c:
    version: "231"
  19200fec012800c5725bf021e47c2ce8baa784f1:
    version: "39"
  22392d51-744e-4699-9a53-763460ff6f42:
    version: "249"
  26869a72-c9e8-40e0-a2cb-55409340c54c:
    version: "112"
  35318da9-3e5f-4f4d-896b-e96a95bf07f8:
    version: "212"
  39420de6b806036821fcdafacd7b2ecf2137e77e:
    version: "21"
  40331cf7-6b02-41ad-9277-b275378c5e78:
    version: "255.6"
  4491ef460d717fec1bef687196299e71663243f7:
    version: "23"
  4744a426-ce86-4bd6-8f5a-0cffc681743d:
    version: "111"
  4987d9d7-ff36-40ca-b01b-fab5eecb791f:
    version: "255.12"
  57966e1cfe9fc55f66933db1de64ba5efc52dbf7:
    version: "12"
  60589dad-345b-4126-87c9-cce546bb52dd:
    version: "176"
  63710ebfa1a5424130d07ce5e621cc91691aa878:
    version: "43"
  66055d1cd1283d54927f0cfe2326dcdb0df02141:
    version: "85"
  81323ee570aac0a8974bb83ea2cb57b40f43b0d4:
    version: "67"
  94066ea7-4d83-4e1b-9fcf-e0fed64eb769:
    version: "233"
  340877edee9f64e8fd9b6897cbe5725fa0b9c244:
    version: "14"
  0360627f-9013-4438-92fa-d65431142bdf:
    version: "156"
  9005037e0772d33a692945597112fa54d542f69d:
    version: "17"
  959408ddcf0a6130e717b3ed4cf2a9a2bdce8604:
    version: "15"
  0990824d-224b-4b71-a9dd-d88231ce4a6f:
    version: "256.9"
  09277055-30c0-4e67-a114-c096c0ee9904:
    version: "164"
  9828318f-1a40-43ed-91d9-22546972d3b6:
    version: "256.2"
  25453110-2f68-490b-ac11-821d518379e0:
    version: "95"
  38053593ac5f3a5032d3e79f153bc286f3a91e50:
    version: "72"
  71305490f6ffab6c14983d0f45f82fd28d3b7144:
    version: "89"
  88231922-cb9a-47bb-a74d-0d1f466afc3a:
    version: "222"
  186178684f3ac482f8cc0c99342376d691de64e1:
    version: "82"
  a3bcfc74-e2bc-4e47-a674-146879b32e60:
    version: "192"
  a5fbc3ce4e3885fd5e1ca0136272afbb873e7161:
    version: "31"
  a6ecc248-0bb1-4365-a657-d950a8536c5e:
    version: "206"
  a06f534a-cf6d-4249-823e-4816de40c395:
    version: "113"
  a11c17b3-ad57-46d8-b782-849e70b5a444:
    version: "102"
  a93bb59a-ca40-4b26-b6f3-45e205312fce:
    version: "137"
  a484d99e74fc3f90d926663a32057e600af7a51f:
    version: "35"
  a0863b1f-4d75-49ea-a8df-ac288e49f4a8:
    version: "100"
  aa6ead50-f2a2-4a78-9b05-3a1e52527c64:
    version: "118"
  aa98b942ac6cef30d7308426fa9b2d0edafd921a:
    version: "3"
  ab0a77d8-b553-4cb2-b653-2c0ced83e106:
    version: "94"
  ab280975f4e4a942c60437846730eeddd9028317:
    version: "92"
  aba487a20d8619d87a74957dc9d988fbcca6064d:
    version: "71"
  ac386c8c-614c-4404-933b-36c2d57e860f:
    version: "252"
  aca1c7ca826b6fc6d945a9d2fc665a40734193a3:
    version: "63"
  acb2b04843353e96041cc588d26016d555823ebf:
    version: "27"
  ad9371f3ba0d55ff061d75a443c44edcc33c5b23:
    version: "28"
  adb10e2a94fc8c26e17fb47f28794ac0a642754e:
    version: "88"
  adb789f6e29f4b70b6dea7fef01144d2d3199120:
    version: "32"
  adb1272a-924c-45ab-a902-bd0f90eca4a1:
    version: "117"
  adcad85a-3cfd-4df9-8a21-bd8525fddaca:
    version: "199"
  addb77c6-4e8d-411f-ab13-b8556faf6464:
    version: "224"
  ae2383a3-4e31-4935-bc5e-fb448ca19a4e:
    version: "144"
  af08a25c725515de237cfd92e4d556fc3a647b56:
    version: "76"
  af749790-d3a7-43cd-9dbc-6d1feaacf199:
    version: "110"
  afed33da-19ac-43f9-ac81-b34716f11f49:
    version: "188"
  b0b251bf-5118-445d-8683-6c5f993ee825:
    version: "216"
  b1af8426-6140-469e-afd1-0ef5860f8857:
    version: "213"
  b1e8e307-778b-498c-9c5f-6d61fd7e8271:
    version: "244"
  b7a3954a-44ec-4267-9846-a85c2d6cee85:
    version: "191"
  b09a26d9-2ef2-4032-918a-d91f05264420:
    version: "181"
  b962c80a-c270-48b6-87c3-fcda9ca34919:
    version: "205"
  b412118f99ad8284b7ba77efc06dbce2dd6d50f6:
    version: "42"
  b9804136-798f-44b7-9b66-434423390161:
    version: "126"
  ba0be057-0474-43f2-6df6-ad3d0b8904fc:
    version: "259"
  ba88d357-e22d-4f4b-9ad0-0707ae477220:
    version: "230"
  ba484e77-1db4-4110-98e3-36967aee4e41:
    version: "172"
  bbfad792cb93660fc7ea7f1020bd6a603f2ed5a3:
    version: "7"
  bc0ff53f-67d5-47d7-872a-d649bc3d9d2a:
    version: "109"
  bde951be-4e44-42fc-6959-208e3b9d634e:
    version: "260.1"
  bdeb4342-53c5-4474-8d17-a0095ef8b680:
    version: "119"
  be6f4858-1fd0-42cc-a251-8fc07f4b80b5:
    version: "197"
  bf1fe13f-7077-405e-b69f-eb62412d2284:
    version: "120"
  bf9678dc-ee9c-4811-a329-223f481411a8:
    version: "103"
  c7a3396f0f19e4433921c3bc9cea01b097b2aec9:
    version: "52"
  c31b8fec04fdf27f8e7a3a361ae0019bf4d0c3b4:
    version: "1"
  c36ae721-65b6-4a7c-b06d-f1218a5395c0:
    version: "106"
  c80ca2f7cfa2a0476e6177abe3dc5c7d09ce154f:
    version: "55"
  c678afe7-7d7f-49cf-9d40-c1e8d8117e71:
    version: "255.4"
  c911d85f806a1b70bd4735c65c3375bd48eb0191:
    version: "66"
  c943e343-214b-408a-9dde-a7188ef2a888:
    version: "142"
  c983fc6800d69b696b23a8b4cc5d29b2c6f07c12:
    version: "77"
  c349618d-11e2-4cdb-9286-8da52620951d:
    version: "158"
  cb987be0-807c-484b-a86a-c28c5af535cc:
    version: "98"
  ce409cc0-f4ff-44eb-ac67-ad22b0a3100d:
    version: "155"
  cec65c68-e31d-4f1b-a9b9-94e0c6ffec64:
    version: "165"
  cf05fee0-79f1-4635-b45d-46a324968d16:
    version: "236"
  cf832f81-b7e9-4a20-82da-27df449e0484:
    version: "256.1"
  cfc286d5-b82d-44af-bd7f-42878150dcff:
    version: "168"
  d00fe303-49ae-4bd5-bdb0-ff10d28f19e7:
    version: "257.15"
  d1dc9119-0910-4d1a-a0fe-8b62129bef39:
    version: "115"
  d5cfd8f9d0c72587e65f7dd498dde62a8794dce6:
    version: "29"
  d78b3b54-b41d-409c-a63b-0af31edc1c60:
    version: "170"
  d7732540-808f-498a-ad97-c826d7db9e7b:
    version: "254"
  d7e497f9-e620-4e47-4065-9a71c0e8b399:
    version: "258"
  d29f79c32afe72544d59598cca3c0e65fbad9db5:
    version: "4"
  d30a5097-c32d-40e9-a402-0e527dd8fea1:
    version: "161"
  d82af44e2c103a8bc7b20d0b8ccc88a266750b3d:
    version: "33"
  d8ba86a3-c25c-4c57-a45a-4587a51b6cda:
    version: "182"
  d514d2325bacf762d34a0642bb7cacd4bd61cfab:
    version: "40"
  d673a62319f43c39b37db7062dec46db2cb66dff:
    version: "44"
  d67965fd-91ac-4551-b091-952c26dd8497:
    version: "234"
  daab772b-3563-4d21-5b39-b1a6acc6fff6:
    version: "260"
  db34bbbb-e923-4e4f-92ec-04a7b94b9886:
    version: "134"
  dc94014e-80a0-41d7-a241-67e0dfee5053:
    version: "153"
  def2b40d-ecf3-470d-8751-773d1cd629d0:
    version: "219"
  df424425-71d3-4df1-b6cd-17cc5c5bff00:
    version: "178"
  dfad3f59-fe62-45d1-9d10-e200a69ebc68:
    version: "209"
  e2c8fc6e-84a0-4d36-b2d5-354b4b2bac91:
    version: "150"
  e38055cd-d428-4d1c-8a3f-a3e983bb5af7:
    version: "253"
  e3a8008cfe47ede4c046c31d5500f9a1a08b002f:
    version: "41"
  e6ca9318-87d9-4ed0-9e99-7232607e253a:
    version: "162"
  e23c7fca-8f19-463f-b479-5e04049ffaa6:
    version: "255.2"
  e45c2bf6dbbddfab0fb29c24c0acf693aec68a1a:
    version: "65"
  e67dc9d432c433516868b216c18b1295f2d29b27:
    version: "53"
  e802a7e147e7fd70e7484fcbaaa65c4f85cd2191:
    version: "45"
  e87ff8de-160c-4096-a29e-753edb103487:
    version: "242"
  e8cc700e-a776-40d4-86a9-60a59eafebfe:
    version: "169"
  e96bb4c6be3a00b24777f57995d3215ff9c92093:
    version: "49"
  e703b2c4-f18e-41e5-836c-59689f88394a:
    version: "220"
  ea25fd55-ae57-4071-9d25-fcd832f38680:
    version: "185"
  ea92a4e7b0bccb29526e9861fe6c77eb5db79062:
    version: "69"
  ea1248de-4029-4c4a-8dae-c36f7c5d307a:
    version: "240"
  eb4a3116-29cb-47b6-99ad-3fa06388eea2:
    version: "122"
  ec126ab0-3eab-453a-9891-2ce83b402dd1:
    version: "251"
  ecb7f48d4c97aff8fc7ce0ccde7ea3efc4c987bb:
    version: "51"
  ece39b52-2478-4d40-9049-e9931663b98e:
    version: "193"
  f3adf894ee76ef6f5c0f7591abb63eaccf06eeea:
    version: "48"
  f3fbedf7-9e2c-4741-abca-fd0a1f7904ab:
    version: "135"
  f6e9386b-bc55-473a-9112-834dbaef0d8b:
    version: "143"
  f7a6d2d5-d843-49a8-9ab5-2d6c5c3a6eb4:
    version: "128"
  f44d6266-fc52-4553-9c63-ce6c2797a022:
    version: "177"
  f65f663c-613c-459f-8338-86d4bdc608c3:
    version: "204"
  f80bc7bd-6234-4246-bc6b-aebeda396176:
    version: "257.9"
  f716dfbc-b698-432e-a9a4-6b094ad82d99:
    version: "174"
  f7442b1f-8a86-467a-8fcb-84b73b11afde:
    version: "201"
  f837286a-38b0-4212-a2d3-8ab18ca8ac5b:
    version: "131"
  fa6df185-8f2a-4a0b-93e5-cf998470c1e9:
    version: "198"
  facaf6f4-8a6e-413e-9cad-34c703ca4880:
    version: "123"
  fc6ec52c08857bc55de73bb81c0016d901644fad:
    version: "56"
  fd075e95-9646-4eb0-81e3-1842bb82e0e5:
    version: "146"
format-version: "2"
`
