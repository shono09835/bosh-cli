package cmd_test

import (
	"errors"

	"github.com/cppforlife/go-patch/patch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shono09835/bosh-cli/v7/cmd"
	fakecmd "github.com/shono09835/bosh-cli/v7/cmd/cmdfakes"
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	fakedir "github.com/shono09835/bosh-cli/v7/director/directorfakes"
	boshtpl "github.com/shono09835/bosh-cli/v7/director/template"
	fakeui "github.com/shono09835/bosh-cli/v7/ui/fakes"
)

var _ = Describe("UpdateRuntimeConfigCmd", func() {
	var (
		ui              *fakeui.FakeUI
		director        *fakedir.FakeDirector
		releaseUploader *fakecmd.FakeReleaseUploader
		command         UpdateRuntimeConfigCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		director = &fakedir.FakeDirector{}
		releaseUploader = &fakecmd.FakeReleaseUploader{
			UploadReleasesStub: func(bytes []byte) ([]byte, error) { return bytes, nil },
		}
		command = NewUpdateRuntimeConfigCmd(ui, director, releaseUploader)
	})

	Describe("Run", func() {
		var (
			opts UpdateRuntimeConfigOpts
		)

		BeforeEach(func() {
			opts = UpdateRuntimeConfigOpts{
				Args: UpdateRuntimeConfigArgs{
					RuntimeConfig: FileBytesArg{Bytes: []byte("runtime: config")},
				},
				Name: "angry-smurf",
			}
		})

		act := func() error { return command.Run(opts) }

		It("updates runtime config", func() {
			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(1))

			name, bytes := director.UpdateRuntimeConfigArgsForCall(0)
			Expect(name).To(Equal("angry-smurf"))
			Expect(bytes).To(Equal([]byte("runtime: config\n")))
		})

		It("updates templated runtime config", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte("name1: ((name1))\nname2: ((name2))"),
			}

			opts.VarKVs = []boshtpl.VarKV{
				{Name: "name1", Value: "val1-from-kv"},
			}

			opts.VarsFiles = []boshtpl.VarsFileArg{
				{Vars: boshtpl.StaticVariables(map[string]interface{}{"name1": "val1-from-file"})},
				{Vars: boshtpl.StaticVariables(map[string]interface{}{"name2": "val2-from-file"})},
			}

			opts.OpsFiles = []OpsFileArg{
				{
					Ops: patch.Ops([]patch.Op{
						patch.ReplaceOp{Path: patch.MustNewPointerFromString("/xyz?"), Value: "val"},
					}),
				},
			}

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(1))

			name, bytes := director.UpdateRuntimeConfigArgsForCall(0)
			Expect(name).To(Equal("angry-smurf"))
			Expect(bytes).To(Equal([]byte("name1: val1-from-kv\nname2: val2-from-file\nxyz: val\n")))
		})

		It("uploads releases provided in the manifest after manifest has been interpolated", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte("before-upload-config: ((key))"),
			}

			opts.VarKVs = []boshtpl.VarKV{
				{Name: "key", Value: "key-val"},
			}

			releaseUploader.UploadReleasesReturns([]byte("after-upload-config"), nil)

			err := act()
			Expect(err).ToNot(HaveOccurred())

			bytes := releaseUploader.UploadReleasesArgsForCall(0)
			Expect(bytes).To(Equal([]byte("before-upload-config: key-val\n")))

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(1))

			name, bytes := director.UpdateRuntimeConfigArgsForCall(0)
			Expect(name).To(Equal("angry-smurf"))
			Expect(bytes).To(Equal([]byte("after-upload-config")))
		})

		It("uploads releases provided in the manifest with fix after manifest has been interpolated", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte("before-upload-config: ((key))"),
			}

			opts.VarKVs = []boshtpl.VarKV{
				{Name: "key", Value: "key-val"},
			}

			opts.FixReleases = true

			releaseUploader.UploadReleasesWithFixReturns([]byte("after-upload-config-with-fix"), nil)

			err := act()
			Expect(err).ToNot(HaveOccurred())

			bytes := releaseUploader.UploadReleasesWithFixArgsForCall(0)
			Expect(bytes).To(Equal([]byte("before-upload-config: key-val\n")))

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(1))

			name, bytes := director.UpdateRuntimeConfigArgsForCall(0)
			Expect(name).To(Equal("angry-smurf"))
			Expect(bytes).To(Equal([]byte("after-upload-config-with-fix")))
		})

		It("returns error and does not deploy if uploading releases fails", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte(`
releases:
- name: capi
  sha1: capi-sha1
  url: https://capi-url
  version: 1+capi
`),
			}

			releaseUploader.UploadReleasesReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(0))
		})

		It("does not update if confirmation is rejected", func() {
			ui.AskedConfirmationErr = errors.New("stop")

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("stop"))

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(0))
		})

		It("returns error if updating failed", func() {
			director.UpdateRuntimeConfigReturns(errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("returns an error if diffing failed", func() {
			director.DiffRuntimeConfigReturns(boshdir.ConfigDiff{}, errors.New("Fetching diff result"))

			err := act()
			Expect(err).To(HaveOccurred())
		})

		It("gets the diff from the deployment", func() {
			diff := [][]interface{}{
				[]interface{}{"some line that stayed", nil},
				[]interface{}{"some line that was added", "added"},
				[]interface{}{"some line that was removed", "removed"},
			}

			expectedDiff := boshdir.NewConfigDiff(diff)
			director.DiffRuntimeConfigReturns(expectedDiff, nil)
			err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(director.DiffRuntimeConfigCallCount()).To(Equal(1))
			Expect(ui.Said).To(ContainElement("  some line that stayed\n"))
			Expect(ui.Said).To(ContainElement("+ some line that was added\n"))
			Expect(ui.Said).To(ContainElement("- some line that was removed\n"))
		})

		Context("when NoRedact option is passed", func() {
			BeforeEach(func() {
				opts = UpdateRuntimeConfigOpts{
					Args: UpdateRuntimeConfigArgs{
						RuntimeConfig: FileBytesArg{Bytes: []byte("runtime: config")},
					},
					Name:     "angry-smurf",
					NoRedact: true,
				}
			})

			It("adds redact to api call", func() {
				director.DiffRuntimeConfigReturns(boshdir.NewConfigDiff([][]interface{}{}), nil)
				err := act()
				Expect(err).ToNot(HaveOccurred())
				_, _, noRedact := director.DiffRuntimeConfigArgsForCall(0)
				Expect(noRedact).To(Equal(true))
			})
		})
	})
})
