package go_samples

import (
	"fmt"
	"net"
)

type InstanceDetail interface {
	GetInstanceIp()
}

type instanceDetail struct {
}

func GetInstanceDetail() InstanceDetail {
	return &instanceDetail{}
}

func (instanceDetail) GetInstanceIp() {
	ifaces, _ := net.Interfaces()
	fmt.Printf("Local IP: %v\n", ifaces)
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Printf("Local IP: %v\n", ip.To4().DefaultMask())
		}
	}
}
