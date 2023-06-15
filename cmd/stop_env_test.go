package cmd_test

import (
	bicmd "github.com/shono09835/bosh-cli/v7/cmd"
	. "github.com/shono09835/bosh-cli/v7/cmd/opts"
	"github.com/cppforlife/go-patch/patch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mock_cmd "github.com/shono09835/bosh-cli/v7/cmd/mocks"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	"github.com/golang/mock/gomock"

	boshtpl "github.com/shono09835/bosh-cli/v7/director/template"
	fakebiui "github.com/shono09835/bosh-cli/v7/ui/fakes"
	fakeui "github.com/shono09835/bosh-cli/v7/ui/fakes"
)

var _ = Describe("StopEnvCmd", func() {
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Run", func() {
		var (
			mockDeploymentStateManager *mock_cmd.MockDeploymentStateManager
			fs                         *fakesys.FakeFileSystem

			fakeUI                 *fakeui.FakeUI
			fakeStage              *fakebiui.FakeStage
			deploymentManifestPath = "/deployment-dir/fake-deployment-manifest.yml"
			statePath              string
			skipDrain              bool
		)

		var newStopEnvCmd = func() *bicmd.StopEnvCmd {
			doGetFunc := func(manifestPath string, statePath_ string, vars boshtpl.Variables, op patch.Op) bicmd.DeploymentStateManager {
				Expect(manifestPath).To(Equal(deploymentManifestPath))
				Expect(vars).To(Equal(boshtpl.NewMultiVars([]boshtpl.Variables{boshtpl.StaticVariables{"key": "value"}})))
				Expect(op).To(Equal(patch.Ops{patch.ErrOp{}}))
				statePath = statePath_
				return mockDeploymentStateManager
			}

			return bicmd.NewStopEnvCmd(fakeUI, doGetFunc)
		}

		var writeDeploymentManifest = func() {
			err := fs.WriteFileString(deploymentManifestPath, `---manifest-content`)
			Expect(err).ToNot(HaveOccurred())
		}

		BeforeEach(func() {
			mockDeploymentStateManager = mock_cmd.NewMockDeploymentStateManager(mockCtrl)
			fs = fakesys.NewFakeFileSystem()
			fs.EnableStrictTempRootBehavior()
			fakeUI = &fakeui.FakeUI{}
			writeDeploymentManifest()
			skipDrain = false
		})

		Context("when skip drain is specified", func() {
			It("gets passed to StopDeployment", func() {
				skipDrain = true
				mockDeploymentStateManager.EXPECT().StopDeployment(skipDrain, fakeStage).Return(nil)
				err := newStopEnvCmd().Run(fakeStage, StopEnvOpts{
					Args: StartStopEnvArgs{
						Manifest: FileBytesWithPathArg{Path: deploymentManifestPath},
					},
					SkipDrain: skipDrain,
					VarFlags: VarFlags{
						VarKVs: []boshtpl.VarKV{{Name: "key", Value: "value"}},
					},
					OpsFlags: OpsFlags{
						OpsFiles: []OpsFileArg{
							{Ops: []patch.Op{patch.ErrOp{}}},
						},
					},
				})
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("state path is NOT specified", func() {
			It("sends the manifest on to the StopDeployment", func() {
				mockDeploymentStateManager.EXPECT().StopDeployment(skipDrain, fakeStage).Return(nil)
				err := newStopEnvCmd().Run(fakeStage, StopEnvOpts{
					Args: StartStopEnvArgs{
						Manifest: FileBytesWithPathArg{Path: deploymentManifestPath},
					},
					SkipDrain: skipDrain,
					VarFlags: VarFlags{
						VarKVs: []boshtpl.VarKV{{Name: "key", Value: "value"}},
					},
					OpsFlags: OpsFlags{
						OpsFiles: []OpsFileArg{
							{Ops: []patch.Op{patch.ErrOp{}}},
						},
					},
				})
				Expect(err).ToNot(HaveOccurred())

				Expect(statePath).To(Equal(""))
			})
		})

		Context("state path is specified", func() {
			It("sends the manifest on to the StopDeployment", func() {
				mockDeploymentStateManager.EXPECT().StopDeployment(skipDrain, fakeStage).Return(nil)
				err := newStopEnvCmd().Run(fakeStage, StopEnvOpts{
					StatePath: "/new/state/file/path/state.json",
					SkipDrain: skipDrain,
					Args: StartStopEnvArgs{
						Manifest: FileBytesWithPathArg{Path: deploymentManifestPath},
					},
					VarFlags: VarFlags{
						VarKVs: []boshtpl.VarKV{{Name: "key", Value: "value"}},
					},
					OpsFlags: OpsFlags{
						OpsFiles: []OpsFileArg{
							{Ops: []patch.Op{patch.ErrOp{}}},
						},
					},
				})
				Expect(err).ToNot(HaveOccurred())

				Expect(statePath).To(Equal("/new/state/file/path/state.json"))
			})
		})

		Context("when the deployment state changer returns an error", func() {
			It("sends the manifest on to the StopDeployment", func() {
				err := bosherr.Error("boom")
				mockDeploymentStateManager.EXPECT().StopDeployment(skipDrain, fakeStage).Return(err)
				returnedErr := newStopEnvCmd().Run(fakeStage, StopEnvOpts{
					Args: StartStopEnvArgs{
						Manifest: FileBytesWithPathArg{Path: deploymentManifestPath},
					},
					SkipDrain: skipDrain,
					VarFlags: VarFlags{
						VarKVs: []boshtpl.VarKV{{Name: "key", Value: "value"}},
					},
					OpsFlags: OpsFlags{
						OpsFiles: []OpsFileArg{
							{Ops: []patch.Op{patch.ErrOp{}}},
						},
					},
				})
				Expect(returnedErr).To(Equal(err))
			})
		})
	})
})
