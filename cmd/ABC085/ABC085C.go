package main

import (
	"fmt"
)

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	for x := 0; x <= n; x++ {
		for y := 0; y <= n-x; y++ {
			z := n - x - y
			total := 10000*x + 5000*y + 1000*z
			if total == t {
				fmt.Println(x, y, z)
				goto Z
			}
		}
	}
	fmt.Println("-1 -1 -1")
Z:
}
