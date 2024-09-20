package netx

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"
)

func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsIPv4(ip string) bool {
	if index := strings.IndexByte(ip, '.'); index > 0 {
		return IsIP(ip)
	}
	return false
}

func IsIPv6(ip string) bool {
	if index := strings.IndexByte(ip, ':'); index >= 0 {
		return IsIP(ip)
	}
	return false
}

func CIDRContainsIP(cidr, ip string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	parsedIP := net.ParseIP(ip)
	return ipnet.Contains(parsedIP)
}

func IsPrivateIP(s string) (isIP bool, isPrivate bool) {
	ip := net.ParseIP(s)
	if ip == nil {
		return false, false
	}
	return true, ip.IsPrivate()
}

func IPExpanded(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return ""
	}
	if ip.To4() != nil {
		return ipStr
	}
	v6, _ := netip.ParseAddr(ipStr)
	return v6.StringExpanded()
}

func IsInvalidPort(s string) bool {
	port, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return port > 0 && port < 65536
}

func IPContains(ips string, target string) bool {
	parsedIP := net.ParseIP(target)
	if parsedIP == nil {
		return false
	}

	if !strings.Contains(ips, "/") {
		// ips is ipv4 or ipv6
		return ips == target
	}

	// ips is cidr
	_, ipNet, err := net.ParseCIDR(ips)
	if err != nil {
		return false
	}

	return ipNet.Contains(parsedIP)
}

// CIDRToIPRange return a list of a CIDR
func CIDRToIPRange(s string) ([]string, error) {
	parsedIP, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}

	inc := func(ip net.IP) {
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] > 0 {
				break
			}
		}
	}

	start, end := ipNet.Mask.Size()
	if start < 1 || end > 128 {
		return nil, fmt.Errorf("invalid CIDR address: %s", s)
	}

	ips := make([]string, 0, 2<<(end-start))
	for ip := parsedIP.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips, nil
}

func IPv4ToLong(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func IPv4StrToLong(ip string) uint32 {
	return IPv4ToLong(net.ParseIP(ip))
}

func LongToIP(longIP uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, longIP)
	return ip
}

func LongToIPv4Str(longIP uint32) string {
	return LongToIP(longIP).String()
}
