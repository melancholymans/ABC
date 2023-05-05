package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	count := 0
	for {
		for i := 0; i < n; i++ {
			if a[i]%2 != 0 {
				goto M
			}
		}
		for i := 0; i < n; i++ {
			a[i] = a[i] / 2
		}
		count += 1
	}
M:
	fmt.Println(count)
}
