package hggw

import (
	"net"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			ip, ok := addr.(*net.IPNet)
			if ok && ip.IP.IsGlobalUnicast() {
				return ip.IP.String()
			}
		}
	}
	return "localhost"
}
