package main

import (
	"fmt"

	"github.com/leiyu321/netlink"
)

func main() {
	addr := netlink.ParseAddr("202.114.6.250/23")
	fmt.Println(netlink.IndexByAddr(addr))
}
