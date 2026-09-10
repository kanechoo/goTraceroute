package socket

import (
	"net"
	"time"
)

// Socket is an abstraction over raw TCP, UDP or ICMP sockets supporting IPv4 and IPv6
type Socket interface {
	// Bind assigns the socket to a local IP address and port
	Bind(addr net.IP, port int) error

	// SendTo sends data to the specified IP address and port
	SendTo(data []byte, flags int, addr net.IP, port int) error

	// RecvFrom reads data from the socket returning number of bytes, source IP, and error if any
	RecvFrom(buf []byte, flags int) (n int, fromAddr net.IP, err error)

	// SetSockOptInt sets a socket option with an integer value
	SetSockOptInt(level, opt, value int) error

	// SetSockOptTimeval sets a socket option with a timeval duration
	SetSockOptTimeval(level, opt int, tv *time.Duration) error

	// BindToDevice pins the socket's egress interface to ifaceName.
	// Unsupported on some platforms; returns an error in that case.
	BindToDevice(ifaceName string) error

	// Close closes the socket and releases resources
	Close() error
}

// SetBindToDeviceFD pins an already-created socket file descriptor to the
// named network interface. It is the low-level primitive shared by the
// Socket implementations and by dialers that need to pick a source address
// on a specific interface.
func SetBindToDeviceFD(fd uintptr, af int, ifaceName string) error {
	return bindToDeviceFD(fd, af, ifaceName)
}

func NewSocket(family, typ, proto int) (Socket, error) {
	return newSocket(family, typ, proto)
}
