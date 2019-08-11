package sshproxy

import (
	"io/ioutil"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	ServerUser    = "user"                    // 接続ユーザー
	ServerHost    = "192.0.2.1"               // 接続先ホスト
	ServerPort    = "22"                      // 接続先ポート
	ServerKeyFile = "/Users/user/.ssh/id_rsa" // 鍵ファイル
)

func LoadPrivateKey(file string) (ssh.AuthMethod, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(key), nil
}

func Connect() (*ssh.Client, error) {
	authByKey, err := LoadPrivateKey(ServerKeyFile)
	if err != nil {
		return nil, err
	}

	conf := &ssh.ClientConfig{
		User: ServerUser,
		Auth: []ssh.AuthMethod{
			authByKey,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // サンプルなので無視
		Timeout:         10 * time.Second,
	}

	conn, err := ssh.Dial("tcp", net.JoinHostPort(ServerHost, ServerPort), conf)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
