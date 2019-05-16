package main

import (
	"fmt"
	"github.com/xinxindu/ip_tool"
)

func main() {

	fmt.Println(ip_tool.IsIp("188.255.123.2"))
	isInner, _ := ip_tool.IsInnerNet("192.168.201.201")
	fmt.Println(isInner)
}
