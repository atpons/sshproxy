package main

import "github.com/atpons/sshproxy"

func main() {
	// sshproxy.Start(sshproxy.NewServer())

	conn, err := sshproxy.Connect()
	defer conn.Close()
	if err != nil {
		return
	}

	sshproxy.Start(sshproxy.NewServerOverSSH(conn))
}
