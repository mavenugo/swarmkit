// +build windows

package xnet

import (
	"net"
	"time"

	"github.com/Microsoft/go-winio"
)

// Listen Wrapper for net.Listen and Windows named pipe
func Listen(proto string, addr string) (net.Listener, error) {
	// Ok this explanation has to go somewhere, so it's going here
	// as of the commit that added this comment (use `git blame`) there is an
	// issue with the way we do listen. basically, `unix` is hard-coded as the
	// protocol for the control api, and `tcp` as the remote api. on windows,
	// we need to use named pipes `npipe` as the control api instead.
	// so, here, we just use a hack. if the protocol is says unix but we're on
	// windows, just use a named pipe anyway. problem changed.
	if proto == "unix" {
		return winio.ListenPipe(addr, "")
	}
	return net.Listen(proto, addr)
}

// DialTimeout Wrapper for net.DialTimeout and Windows named pipe (winio.DialPipe)
func DialTimeout(proto string, addr string, timeout time.Duration) (net.Conn, error) {
	if proto == "npipe" {
		return winio.DialPipe(addr, &timeout)
	}
	return net.DialTimeout(proto, addr, timeout)
}
