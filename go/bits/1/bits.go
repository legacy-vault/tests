// bits.go

package main

import (
	"fmt"
	"math/bits"
)

func main() {

	var a, b, c, d uint8
	var e, f, g, h, i, j, k, l uint8

	a = 1
	b = 2
	c = 128
	d = 64

	fmt.Println(bits.LeadingZeros8(a), bits.LeadingZeros8(b), bits.LeadingZeros8(c), bits.LeadingZeros8(d))
	fmt.Println(bits.Len8(a), bits.Len8(b), bits.Len8(c), bits.Len8(d))

	e = 5
	f = bits.Reverse8(e)
	g = bits.RotateLeft8(e, 1)
	h = bits.RotateLeft8(e, 3)
	i = bits.RotateLeft8(e, -1)
	j = bits.RotateLeft8(e, -3)
	k = e << 1
	l = e >> 1

	fmt.Println(bits.OnesCount8(e))
	fmt.Printf("E = %v = [%b]\r\n", e, e)
	fmt.Printf("F = %v = [%b]\r\n", f, f)
	fmt.Printf("G = %v = [%b]\r\n", g, g)
	fmt.Printf("H = %v = [%b]\r\n", h, h)
	fmt.Printf("I = %v = [%b]\r\n", i, i)
	fmt.Printf("J = %v = [%b]\r\n", j, j)
	fmt.Printf("K = %v = [%b]\r\n", k, k)
	fmt.Printf("L = %v = [%b]\r\n", l, l)

	fmt.Println(bits.TrailingZeros8(a), bits.TrailingZeros8(b), bits.TrailingZeros8(c), bits.TrailingZeros8(d))
}
