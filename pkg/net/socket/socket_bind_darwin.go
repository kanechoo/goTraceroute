//go:build darwin

package socket

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"
)

// bindToDeviceFD pins the socket referred to by fd to the network interface
// named ifaceName.
//
// Darwin has no SO_BINDTODEVICE. The equivalent is IP_BOUND_IF (IPv4) /
// IPV6_BOUND_IF (IPv6), both of which take the interface index. This affects
// both the source-address lookup and the egress interface chosen by the
// routing table, which lets callers bypass a VPN/tunnel that owns the default
// route.
func bindToDeviceFD(fd uintptr, af int, ifaceName string) error {
	if ifaceName == "" {
		return nil
	}

	ifi, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return fmt.Errorf("bind to interface %q: %w", ifaceName, err)
	}

	var level, opt int
	switch af {
	case AF_INET:
		level, opt = unix.IPPROTO_IP, unix.IP_BOUND_IF
	case AF_INET6:
		level, opt = unix.IPPROTO_IPV6, unix.IPV6_BOUND_IF
	default:
		return fmt.Errorf("bind to interface %q: unsupported address family %d", ifaceName, af)
	}

	if err := unix.SetsockoptInt(int(fd), level, opt, ifi.Index); err != nil {
		return fmt.Errorf("bind to interface %q: %w", ifaceName, err)
	}
	return nil
}
