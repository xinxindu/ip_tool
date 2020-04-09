package ip_tool

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

/*
	判断是否内网IP
*/
func IsInnerNet(ip string) (bool, error) {
	if !IsIp(ip) {
		return false, errors.New("parameter is not ip")
	}

	if ip == "127.0.0.1" {
		return true, nil
	}

	sections := strings.Split(ip, ".")
	switch sections[0] {
	case "10":
		return true, nil
	case "172":
		num, _ := strconv.Atoi(sections[1])
		if num >= 16 && num <= 31 {
			return true, nil
		}
		return false, nil
	case "192":
		if sections[1] == "168" {
			return true, nil
		}
		return false, nil
	default:
		return false, nil
	}
}

/*
	判断是否是IP
*/
func IsIp(ipStr string) bool {
	return net.ParseIP(ipStr) != nil
}