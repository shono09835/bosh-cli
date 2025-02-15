package sshtunnel

import (
	"fmt"
	"io"
	"net"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	boshssh "github.com/shono09835/bosh-cli/v7/ssh"
)

type SSHTunnel interface {
	Start(chan<- error, chan<- error)
}

type sshTunnel struct {
	client boshssh.Client

	localForwardPort  int
	remoteForwardPort int

	remoteListener net.Listener

	logTag string
	logger boshlog.Logger
}

func (s *sshTunnel) Start(readyErrCh chan<- error, errCh chan<- error) {
	err := s.client.Start()
	if err != nil {
		readyErrCh <- bosherr.WrapError(err, "Starting SSH tunnel")
		return
	}

	remoteListenAddr := fmt.Sprintf("127.0.0.1:%d", s.remoteForwardPort)
	s.logger.Debug(s.logTag, "Listening on remote server %s", remoteListenAddr)
	s.remoteListener, err = s.client.Listen("tcp", remoteListenAddr)
	if err != nil {
		readyErrCh <- bosherr.WrapError(err, "Listening on remote server")
		return
	}

	readyErrCh <- nil

	for {
		remoteConn, err := s.remoteListener.Accept()
		s.logger.Debug(s.logTag, "Received connection")
		if err != nil {
			errCh <- bosherr.WrapError(err, "Accepting connection on remote server")
			return
		}

		defer func() {
			if err = remoteConn.Close(); err != nil {
				s.logger.Warn(s.logTag, "Failed to close remote listener connection: %s", err.Error())
			}
		}()

		s.logger.Debug(s.logTag, "Dialing local server")
		localDialAddr := fmt.Sprintf("127.0.0.1:%d", s.localForwardPort)
		localConn, err := net.Dial("tcp", localDialAddr)
		if err != nil {
			errCh <- bosherr.WrapError(err, "Dialing local server")
			return
		}

		go func(localConn, remoteConn net.Conn) {
			defer func() {
				if err = localConn.Close(); err != nil {
					s.logger.Warn(s.logTag, "Failed to close local dial connection: %s", err.Error())
				}
			}()

			bytesNum, err := io.Copy(remoteConn, localConn)

			s.logger.Debug(s.logTag, "Copying bytes from local to remote %d", bytesNum)

			if err != nil {
				errCh <- bosherr.WrapError(err, "Copying bytes from local to remote")
			}
		}(localConn, remoteConn)

		go func(localConn, remoteConn net.Conn) {
			defer func() {
				if err = localConn.Close(); err != nil {
					s.logger.Warn(s.logTag, "Failed to close local dial connection: %s", err.Error())
				}
			}()

			bytesNum, err := io.Copy(localConn, remoteConn)

			s.logger.Debug(s.logTag, "Copying bytes from remote to local %d", bytesNum)

			if err != nil {
				errCh <- bosherr.WrapError(err, "Copying bytes from remote to local")
			}
		}(localConn, remoteConn)
	}
}
