package helpers

import (
	"net"
	"strings"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetConnectIp() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, i := range ifaces {
		//只抓取網路卡名稱為"en0", "en1"...
		if strings.Contains(i.Name, "ppp") {
			addrs, err := i.Addrs()
			if err != nil {
				//log.Println("No IP:", err)
				return ""
			}

			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					//log.Println("IPNET")
					ip = v.IP
				case *net.IPAddr:
					//log.Println("IPAddr")
					ip = v.IP
				}

				//這裡會抓取兩種IP，分別是IPv4與IPv6
				if ip[0] == 0 {
					//第一個byte是0為IPv4
					//log.Println("Get device:", i.Name)
					return ip.String()
				}
			}
		}
	}

	return ""
}
