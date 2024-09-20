package netx

import (
	"net"
	"time"
)

func Detect(ip, port string) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
