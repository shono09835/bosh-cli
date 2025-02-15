package ssh_test

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"syscall"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	boshdir "github.com/shono09835/bosh-cli/v7/director"
	. "github.com/shono09835/bosh-cli/v7/ssh"
	fakessh "github.com/shono09835/bosh-cli/v7/ssh/sshfakes"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
	fakeui "github.com/shono09835/bosh-cli/v7/ui/fakes"
)

var _ = Describe("ComboRunner", func() {
	var (
		cmdRunner   *fakesys.FakeCmdRunner
		session     *fakessh.FakeSession
		signalCh    chan<- os.Signal
		writer      Writer
		fs          *fakesys.FakeFileSystem
		ui          *fakeui.FakeUI
		logger      boshlog.Logger
		comboRunner ComboRunner
	)

	BeforeEach(func() {
		cmdRunner = fakesys.NewFakeCmdRunner()

		session = &fakessh.FakeSession{}
		sessFactory := func(_ ConnectionOpts, _ boshdir.SSHResult) Session { return session }

		signalCh = nil
		signalNotifyFunc := func(ch chan<- os.Signal, s ...os.Signal) { signalCh = ch }

		ui = &fakeui.FakeUI{}

		writer = NewStreamingWriter(boshui.NewComboWriter(ui))

		fs = fakesys.NewFakeFileSystem()
		fs.ReturnTempFilesByPrefix = map[string]boshsys.File{
			"ssh-priv-key":    fakesys.NewFakeFile("/tmp/priv-key", fs),
			"ssh-known-hosts": fakesys.NewFakeFile("/tmp/known-hosts", fs),
		}

		logger = boshlog.NewLogger(boshlog.LevelNone)

		comboRunner = NewComboRunner(
			cmdRunner, sessFactory, signalNotifyFunc, writer, fs, ui, logger)
	})

	Describe("Run", func() {
		var (
			connOpts   ConnectionOpts
			result     boshdir.SSHResult
			cmdFactory func(host boshdir.Host, args SSHArgs) boshsys.Command
		)

		BeforeEach(func() {
			connOpts = ConnectionOpts{}
			result = boshdir.SSHResult{}
			cmdFactory = func(host boshdir.Host, args SSHArgs) boshsys.Command {
				return boshsys.Command{Name: "cmd", Args: []string{host.Host}}
			}
		})

		It("returns without error when there are no hosts", func() {
			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).ToNot(HaveOccurred())

			Expect(session.FinishCallCount()).To(Equal(1))
		})

		It("returns without error when there is only one host", func() {
			result.Hosts = []boshdir.Host{{Host: "127.0.0.1"}}

			cmdRunner.AddProcess("cmd 127.0.0.1", &fakesys.FakeProcess{})

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).ToNot(HaveOccurred())

			Expect(session.FinishCallCount()).To(Equal(1))
		})

		It("returns without error when there are multiple hosts", func() {
			result.Hosts = []boshdir.Host{
				{Host: "127.0.0.1"},
				{Host: "127.0.0.2"},
			}

			cmdRunner.AddProcess("cmd 127.0.0.1", &fakesys.FakeProcess{})
			cmdRunner.AddProcess("cmd 127.0.0.2", &fakesys.FakeProcess{})

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).ToNot(HaveOccurred())

			Expect(session.FinishCallCount()).To(Equal(1))
		})

		It("provides ssh arguments to customize cmd", func() {
			cmdFactory = func(host boshdir.Host, args SSHArgs) boshsys.Command {
				optsStr := strings.Join(args.LoginForHost(host), " ")
				var opts []string
				switch {
				case strings.Contains(optsStr, "127.0.0.1"):
					opts = []string{"ip-1"}
				case strings.Contains(optsStr, "127.0.0.2"):
					opts = []string{"ip-2"}
				default:
					panic("Unexpected ssh args")
				}
				return boshsys.Command{Name: "cmd", Args: opts}
			}

			result.Hosts = []boshdir.Host{
				{Host: "127.0.0.1", Username: "user-1"},
				{Host: "127.0.0.2", Username: "user-2"},
			}

			sshArgs := SSHArgs{
				ConnOpts: connOpts,
				Result:   result,

				PrivKeyFile:    fakesys.NewFakeFile("/priv-key", fs),
				KnownHostsFile: fakesys.NewFakeFile("/priv-key", fs),
			}

			session.StartReturns(sshArgs, nil)

			cmdRunner.AddProcess("cmd ip-1", &fakesys.FakeProcess{})
			cmdRunner.AddProcess("cmd ip-2", &fakesys.FakeProcess{})

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).ToNot(HaveOccurred())
		})

		It("writes to ui with a instance prefix", func() {
			result.Hosts = []boshdir.Host{
				{Job: "job1", IndexOrID: "id1", Host: "127.0.0.1"},
				{Job: "job2", IndexOrID: "id2", Host: "127.0.0.2"},
			}

			proc1 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.1", proc1)

			proc2 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.2", proc2)

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).ToNot(HaveOccurred())

			_, err = proc1.Stdout.Write([]byte("stdout1\n"))
			Expect(err).ToNot(HaveOccurred())
			_, err = proc1.Stderr.Write([]byte("stderr1\n"))
			Expect(err).ToNot(HaveOccurred())

			_, err = proc2.Stdout.Write([]byte("stdout2\n"))
			Expect(err).ToNot(HaveOccurred())
			_, err = proc2.Stderr.Write([]byte("stderr2\n"))
			Expect(err).ToNot(HaveOccurred())

			Expect(ui.Blocks).To(Equal([]string{
				"job1/id1: stdout | ", "stdout1", "\n",
				"job1/id1: stderr | ", "stderr1", "\n",
				"job2/id2: stdout | ", "stdout2", "\n",
				"job2/id2: stderr | ", "stderr2", "\n",
			}))
		})

		It("writes to ui with a ? prefix when job name is not known", func() {
			result.Hosts = []boshdir.Host{
				{Job: "", IndexOrID: "id1", Host: "127.0.0.1"},
				{Job: "job2", IndexOrID: "id2", Host: "127.0.0.2"},
			}

			proc1 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.1", proc1)

			proc2 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.2", proc2)

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).ToNot(HaveOccurred())

			_, err = proc1.Stdout.Write([]byte("stdout1\n"))
			Expect(err).ToNot(HaveOccurred())
			_, err = proc2.Stdout.Write([]byte("stdout2\n"))
			Expect(err).ToNot(HaveOccurred())

			_, err = proc1.Stderr.Write([]byte("stderr1\n"))
			Expect(err).ToNot(HaveOccurred())
			_, err = proc2.Stderr.Write([]byte("stderr2\n"))
			Expect(err).ToNot(HaveOccurred())

			Expect(ui.Blocks).To(Equal([]string{
				"?/id1: stdout | ", "stdout1", "\n",
				"job2/id2: stdout | ", "stdout2", "\n",
				"?/id1: stderr | ", "stderr1", "\n",
				"job2/id2: stderr | ", "stderr2", "\n",
			}))
		})

		It("uses provided stdout/stderr if given", func() {
			result.Hosts = []boshdir.Host{
				{Host: "127.0.0.1"},
			}

			proc1 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.1", proc1)

			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")

			cmdFactory = func(host boshdir.Host, args SSHArgs) boshsys.Command {
				return boshsys.Command{
					Name:   "cmd",
					Args:   []string{host.Host},
					Stdout: stdout,
					Stderr: stderr,
				}
			}

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).ToNot(HaveOccurred())

			_, err = proc1.Stdout.Write([]byte("stdout"))
			Expect(err).ToNot(HaveOccurred())
			_, err = proc1.Stderr.Write([]byte("stderr"))
			Expect(err).ToNot(HaveOccurred())

			Expect(stdout.String()).To(Equal("stdout"))
			Expect(stderr.String()).To(Equal("stderr"))
		})

		It("ultimately returns an error if any processes fail to start", func() {
			result.Hosts = []boshdir.Host{
				{Host: "127.0.0.1"},
				{Host: "127.0.0.2"},
			}

			proc1 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.1", proc1)

			proc2 := &fakesys.FakeProcess{
				StartErr: errors.New("fake-err"),
			}
			cmdRunner.AddProcess("cmd 127.0.0.2", proc2)

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(session.FinishCallCount()).To(Equal(1))
		})

		It("ultimately returns an error if any processes fail during execution", func() {
			result.Hosts = []boshdir.Host{
				{Host: "127.0.0.1"},
				{Host: "127.0.0.2"},
			}

			proc1 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.1", proc1)

			proc2 := &fakesys.FakeProcess{
				WaitResult: boshsys.Result{Error: errors.New("fake-err")},
			}
			cmdRunner.AddProcess("cmd 127.0.0.2", proc2)

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(session.FinishCallCount()).To(Equal(1))
		})

		It("includes all errors if any processes fail", func() {
			result.Hosts = []boshdir.Host{
				{Host: "127.0.0.1"},
				{Host: "127.0.0.2"},
				{Host: "127.0.0.3"},
			}

			proc1 := &fakesys.FakeProcess{}
			cmdRunner.AddProcess("cmd 127.0.0.1", proc1)

			proc2 := &fakesys.FakeProcess{
				WaitResult: boshsys.Result{Error: errors.New("fake-err2")},
			}
			cmdRunner.AddProcess("cmd 127.0.0.2", proc2)

			proc3 := &fakesys.FakeProcess{
				StartErr: errors.New("fake-err3"),
			}
			cmdRunner.AddProcess("cmd 127.0.0.3", proc3)

			err := comboRunner.Run(connOpts, result, cmdFactory)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err2"))
			Expect(err.Error()).To(ContainSubstring("fake-err3"))
		})

		Describe("signal handling", func() {
			var errCh chan error

			BeforeEach(func() {
				errCh = make(chan error)

				result.Hosts = []boshdir.Host{
					{Host: "127.0.0.1"},
					{Host: "127.0.0.2"},
					{Host: "127.0.0.3"},
				}

				proc1 := &fakesys.FakeProcess{
					TerminatedNicelyCallBack: func(p *fakesys.FakeProcess) {
						p.WaitCh <- boshsys.Result{}
					},
				}
				cmdRunner.AddProcess("cmd 127.0.0.1", proc1)

				proc2 := &fakesys.FakeProcess{}
				cmdRunner.AddProcess("cmd 127.0.0.2", proc2)

				proc3 := &fakesys.FakeProcess{
					TerminatedNicelyCallBack: func(p *fakesys.FakeProcess) {
						p.WaitCh <- boshsys.Result{Error: errors.New("term-err")}
					},
				}
				cmdRunner.AddProcess("cmd 127.0.0.3", proc3)

				go func() {
					defer GinkgoRecover()
					// Wait for interrupt goroutine to set channel
					errCh <- comboRunner.Run(connOpts, result, cmdFactory)
				}()

				Eventually(func() chan<- os.Signal { return signalCh }).ShouldNot(BeNil())
			})

			It("terminates processes nicely upon interrupt", func() {
				Eventually(signalCh).Should(BeSent(os.Interrupt))

				var err error
				Eventually(errCh).Should(Receive(&err))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("term-err"))
				Expect(session.FinishCallCount()).To(Equal(2))
			})

			It("terminates processes nicely upon sigterm", func() {
				Eventually(signalCh).Should(BeSent(syscall.SIGTERM))

				var err error
				Eventually(errCh).Should(Receive(&err))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("term-err"))
				Expect(session.FinishCallCount()).To(Equal(2))
			})

			It("terminates processes nicely upon sighup", func() {
				Eventually(signalCh).Should(BeSent(syscall.SIGHUP))

				var err error
				Eventually(errCh).Should(Receive(&err))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("term-err"))
				Expect(session.FinishCallCount()).To(Equal(2))
			})
		})
	})
})
