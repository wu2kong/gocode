package main

import (
	"fmt"
	"net"
)

func GetLocalIPv4Address() (string, error) {
	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range interfaces {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet)
			}

		}
	}
	return "", nil
}

func main() {
	iplocal, err := GetLocalIPv4Address()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", iplocal)
}
