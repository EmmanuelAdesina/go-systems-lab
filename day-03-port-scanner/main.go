package main

import (
	"fmt"
	"net"
	"strconv"
)

func ScanPort(host string, port int) (bool, error) {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return false, err
	}

	defer conn.Close()
	return true, nil
}

func main() {
	host := "google.com"
	port := 443

	open, err := ScanPort(host, port)

	if err != nil {
		fmt.Println(err)
		return
	}

	if open {
		fmt.Println("Port is open")
	} else {
		fmt.Println("Port is closed")
	}
}
