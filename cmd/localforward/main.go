package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/atpons/sshproxy"
)

func main() {
	conn, err := sshproxy.Connect()
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: conn.Dial,
		},
	}

	resp, err := client.Get("http://httpbin.org/ip")

	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
