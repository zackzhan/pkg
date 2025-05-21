package ip

import (
	"errors"
	"net"
)

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}

		if ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		} else if ipNet.IP.To16() != nil {
			return ipNet.IP.String(), nil
		}
	}

	return "", errors.New("no suitable local IP address found")
}
