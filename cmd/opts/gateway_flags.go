package opts

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	boshdir "github.com/shono09835/bosh-cli/v7/director"
	boshssh "github.com/shono09835/bosh-cli/v7/ssh"
)

func (f GatewayFlags) AsSSHOpts() (boshdir.SSHOpts, boshssh.ConnectionOpts, error) {
	sshOpts, privKey, err := boshdir.NewSSHOpts(f.UUIDGen)
	if err != nil {
		return boshdir.SSHOpts{}, boshssh.ConnectionOpts{}, bosherr.WrapErrorf(err, "Generating SSH options")
	}

	connOpts := boshssh.ConnectionOpts{
		PrivateKey: privKey,

		GatewayDisable: f.Disable,

		GatewayUsername:       f.Username,
		GatewayHost:           f.Host,
		GatewayPrivateKeyPath: f.PrivateKeyPath,

		SOCKS5Proxy: f.SOCKS5Proxy,
	}

	return sshOpts, connOpts, nil
}
