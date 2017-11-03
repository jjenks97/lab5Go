package main

import (
	"fmt"
	"math"
	"golang.org/x/tour/wc"
	"strings"
)

func Sqrt(x float64) float64 {
	z := x / 2
	y := 0.0
	for math.Abs(y-z) > 1e-5 {
		y = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)

	for _, x := range words{
		wordCount[x]++
	}

	return wordCount
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	fmt.Println(Sqrt(2))
	wc.Test(WordCount)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
