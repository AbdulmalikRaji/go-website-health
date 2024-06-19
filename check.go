package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := fmt.Sprintf("%s:%s", domain, port)
	timeout := time.Duration(time.Second * 5)
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err!= nil {
        return fmt.Sprintf("[DOWN] Error connecting to %s: %v", domain, err)
    }
	defer conn.Close()
	return fmt.Sprintf("[UP] Connected to %s: \n from: %v \n to: %v ", domain, conn.LocalAddr(), conn.RemoteAddr())
}
