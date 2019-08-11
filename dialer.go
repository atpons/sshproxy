package sshproxy

import (
	"net"
	"strconv"

	"github.com/cybozu-go/usocksd/socks"
	"github.com/nbio/contextual"
	"golang.org/x/crypto/ssh"
)

type ForwardDialer struct {
	*ssh.Client
}

func (f ForwardDialer) Dial(r *socks.Request) (net.Conn, error) {
	var addr string
	if len(r.Hostname) > 0 {
		addr = net.JoinHostPort(r.Hostname, strconv.Itoa(r.Port))
	} else {
		addr = net.JoinHostPort(r.IP.String(), strconv.Itoa(r.Port))
	}

	dialer := contextual.Dialer{SimpleDialer: f.Client}

	return dialer.DialContext(r.Context(), "tcp", addr)
}
