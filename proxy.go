package sshproxy

import (
	"net"
	"os"

	"github.com/cybozu-go/usocksd"
	"github.com/cybozu-go/usocksd/socks"
	"github.com/cybozu-go/well"
	"golang.org/x/crypto/ssh"
)

var (
	ProxyConf = &usocksd.Config{
		Incoming: usocksd.IncomingConfig{
			Port:      8080,
			Addresses: []net.IP{net.ParseIP("127.0.0.1")},
		},
	}
)

func Start(s *socks.Server) {
	g := &well.Graceful{
		Listen: func() ([]net.Listener, error) {
			return usocksd.Listeners(ProxyConf)
		},
		Serve: func(lns []net.Listener) {
			for _, ln := range lns {
				s.Serve(ln)
			}
			err := well.Wait()
			if err != nil && !well.IsSignaled(err) {
				os.Exit(1)
			}
		},
		ExitTimeout: 1,
	}
	g.Run()

	err := well.Wait()
	if err != nil && !well.IsSignaled(err) {
		os.Exit(1)
	}
}

func Serve(lns []net.Listener) {
	s := NewServer()
	for _, ln := range lns {
		s.Serve(ln)
	}
	err := well.Wait()
	if err != nil && !well.IsSignaled(err) {
		os.Exit(1)
	}
}

func NewServer() *socks.Server {
	s := usocksd.NewServer(ProxyConf)
	return s
}

func NewServerOverSSH(client *ssh.Client) *socks.Server {
	s := usocksd.NewServer(ProxyConf)
	s.Dialer = ForwardDialer{client}
	return s
}
