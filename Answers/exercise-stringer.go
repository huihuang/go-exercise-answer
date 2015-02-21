/* Exercise: Stringers */

package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (add IPAddr) String() string{
	return string(strconv.Itoa(int(add[0])) + "." + strconv.Itoa(int(add[1])) + "." + strconv.Itoa(int(add[2])) + "." + strconv.Itoa(int(add[3])))
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
