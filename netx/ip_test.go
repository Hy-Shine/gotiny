package netx

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidIP(t *testing.T) {
	cases := []struct {
		in  string
		exp bool
	}{
		{in: "::1", exp: true},
		{in: "127.0.0.1", exp: true},
		{in: "192.168..", exp: false},
		{in: "192.168..1", exp: false},
	}

	for _, v := range cases {
		assert.Equal(t, v.exp, IsIP(v.in))
	}
}

func TestIsValidIPv4(t *testing.T) {
	cases := []struct {
		in  string
		exp bool
	}{
		{in: "::1", exp: false},
		{in: "127.0.0.1", exp: true},
		{in: "192.168.0.", exp: false},
		{in: "192.168..1", exp: false},
		{in: "", exp: false},
		{in: "127.0.0.1.", exp: false},
	}
	for _, v := range cases {
		assert.Equal(t, v.exp, IsIPv4(v.in), fmt.Sprintf("case: %s not pass", v.in))
	}
}

func TestIsValidIPv6(t *testing.T) {
	cases := []struct {
		in  string
		exp bool
	}{
		{in: "::1", exp: true},
		{in: "127.0.0.1", exp: false},
		{in: "1080::8:800:200C:417A", exp: true},
		{in: "", exp: false},
		{in: "FF01::101", exp: true},
		{in: "1080:0:0:0:8:800:200C:417A", exp: true},
	}
	for _, v := range cases {
		assert.Equal(t, v.exp, IsIPv6(v.in), fmt.Sprintf("case: %s NOT PASS", v.in))
	}
}

func TestCIDRContainsIP(t *testing.T) {
	cases := []struct {
		cidr, ip string
		exp      bool
	}{
		{cidr: "192.168.1.0/24", ip: "192.168.1.100", exp: true},
		{cidr: "192.168.1.0/24", ip: "192.168.2.100", exp: false},
		{cidr: "172.17.0.0/16", ip: "172.17.0.10", exp: true},
		{cidr: "10.0.0.0/24", ip: "127.0.0.1", exp: false},
	}
	for _, v := range cases {
		assert.Equal(t, v.exp, CIDRContainsIP(v.cidr, v.ip), fmt.Sprintf("cidr:%s ip:%s NOT PASS", v.cidr, v.ip))
	}
}

func TestToLong(t *testing.T) {
	cases := []struct {
		ip     string
		expect uint32
	}{
		{ip: "192.168.0.1", expect: 3232235521},
		{ip: "10.0.0.0", expect: 167772160},
		{ip: "255.255.255.255", expect: 4294967295},
		{ip: "0.0.0.0", expect: 0},
	}

	for _, c := range cases {
		ip := net.ParseIP(c.ip)
		got := IPv4ToLong(ip)
		if got != c.expect {
			t.Errorf("IPv4ToLong(%s) = %d, expect %d", c.ip, got, c.expect)
		}
	}
}

func TestLongToIP(t *testing.T) {
	longIP := IPv4ToLong(net.ParseIP("192.168.0.1"))
	fmt.Println(LongToIP(longIP))
}

func TestIPv4ToLong(t *testing.T) {
	cases := []struct {
		ip     string
		expect uint32
	}{
		{ip: "192.168.0.1", expect: 3232235521},
		{ip: "10.0.0.0", expect: 167772160},
		{ip: "255.255.255.255", expect: 4294967295},
		{ip: "0.0.0.0", expect: 0},
	}

	for _, c := range cases {
		ip := net.ParseIP(c.ip)
		got := IPv4ToLong(ip)
		if got != c.expect {
			t.Errorf("IPv4ToLong(%s) = %d, expect %d", c.ip, got, c.expect)
		}
	}
}

func BenchmarkIPv4ToLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv4ToLong(net.ParseIP("192.168.0.1"))
	}
}

func TestIPv4StrToLong(t *testing.T) {
	cases := []struct {
		ip     string
		expect uint32
	}{
		{ip: "192.168.0.1", expect: 3232235521},
		{ip: "10.0.0.0", expect: 167772160},
		{ip: "255.255.255.255", expect: 4294967295},
		{ip: "0.0.0.0", expect: 0},
	}

	for _, c := range cases {
		got := IPv4StrToLong(c.ip)
		if got != c.expect {
			t.Errorf("IPv4StrToLong(%s) = %d, expect %d", c.ip, got, c.expect)
		}
	}
}

func BenchmarkIPvStrToLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv4StrToLong("192.168.0.1")
	}
}

func TestLongToIPv4Str(t *testing.T) {
	tests := []struct {
		name string
		ip   uint32
		want string
	}{
		{
			name: "Zero IP",
			ip:   0,
			want: "0.0.0.0",
		},
		{
			name: "Loopback IP",
			ip:   2130706433,
			want: "127.0.0.1",
		},
		{
			name: "Google DNS IP",
			ip:   134744072,
			want: "8.8.8.8",
		},
		{
			name: "Max IP",
			ip:   4294967295,
			want: "255.255.255.255",
		},
		{
			name: "Cloudflare DNS IP",
			ip:   16843009,
			want: "1.1.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongToIPv4Str(tt.ip); got != tt.want {
				t.Errorf("LongToIPv4Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
