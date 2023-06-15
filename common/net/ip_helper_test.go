package net_test

import (
	"net"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	binet "github.com/shono09835/bosh-cli/v7/common/net"
)

var _ = Describe("common.net ip helpers", func() {
	Describe("LastAddress", func() {
		It("Returns the last address in the network range", func() {
			Expect(
				binet.LastAddress(netFor("10.0.0.0/24")),
			).To(Equal(net.ParseIP("10.0.0.255")))

			Expect(
				binet.LastAddress(netFor("10.0.0.20/24")),
			).To(Equal(net.ParseIP("10.0.0.255")))

			Expect(
				binet.LastAddress(netFor("10.1.0.0/24")),
			).To(Equal(net.ParseIP("10.1.0.255")))

			Expect(
				binet.LastAddress(netFor("10.10.0.0/8")),
			).To(Equal(net.ParseIP("10.255.255.255")))

			Expect(
				binet.LastAddress(netFor("2001:db8:1234::/48")),
			).To(Equal(net.ParseIP("2001:db8:1234:ffff:ffff:ffff:ffff:ffff")))
		})
	})
})

func netFor(ipNetString string) *net.IPNet {
	_, ipNet, err := net.ParseCIDR(ipNetString)
	Expect(err).ToNot(HaveOccurred())
	return ipNet
}
