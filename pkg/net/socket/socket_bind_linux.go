//go:build linux

package socket

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// bindToDeviceFD pins the socket referred to by fd to the network interface
// named ifaceName via SO_BINDTODEVICE.
//
// This affects both the source-address lookup and the egress interface chosen
// by the routing table, which lets callers bypass a VPN/tunnel that owns the
// default route. The operation requires CAP_NET_RAW/CAP_NET_ADMIN (root).
func bindToDeviceFD(fd uintptr, af int, ifaceName string) error {
	if ifaceName == "" {
		return nil
	}

	if err := unix.SetsockoptString(int(fd), unix.SOL_SOCKET, unix.SO_BINDTODEVICE, ifaceName); err != nil {
		return fmt.Errorf("bind to interface %q: %w", ifaceName, err)
	}
	return nil
}
