package ip

import "net"

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}

		if ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		} else if ipNet.IP.To16() != nil {
			return ipNet.IP.String()
		}
	}

	return ""
}
