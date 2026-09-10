//go:build !darwin && !linux

package socket

import "fmt"

// bindToDeviceFD is only implemented on Darwin and Linux.
func bindToDeviceFD(fd uintptr, af int, ifaceName string) error {
	if ifaceName == "" {
		return nil
	}
	return fmt.Errorf("binding a socket to interface %q is not supported on this platform", ifaceName)
}
