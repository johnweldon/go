package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var network bool

func init() {
	flag.BoolVar(&network, "m", true, "do network mask calculation")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if network && len(args) == 2 {
		base := net.ParseIP(args[0])
		end := net.ParseIP(args[1])
		pfx, _ := CommonPrefixLength(base, end)
		mask := net.CIDRMask(int(pfx), 32)
		net := net.IPNet{IP: base.Mask(mask), Mask: mask}
		fmt.Fprintf(os.Stdout, "%s\n", net.String())
	} else {
		flag.Usage()
	}
}

func CommonPrefixLength(lhs, rhs net.IP) (uint8, bool) {
	l, r := lhs.To4(), rhs.To4()
	if l == nil || r == nil {
		return 0, false
	}

	c := 0
	for bx := 0; bx < 4; bx++ {
		if l[bx] == r[bx] {
			c = 8 * (bx + 1)
		} else {
			for i := uint8(7); i != 0; i-- {
				if l[bx]>>i == r[bx]>>i {
					c += 1
				} else {
					return uint8(c), true
				}
			}
		}
	}
	return uint8(c), true
}
