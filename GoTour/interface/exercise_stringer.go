// IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

package main

import "fmt"
import "strings"

type IPAddr [4]byte // <---

func (ipAddr IPAddr) String() string {

	var builder strings.Builder 
	for i, b := range ipAddr {
		if i > 0 {
			builder.WriteString(".")
		}
		builder.WriteString(fmt.Sprintf("%d", b)) // formats and returns a string in decimal format without printing it
	}
	return builder.String()
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
