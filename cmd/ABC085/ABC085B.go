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
	m := make(map[int]bool)
	uniq := []int{}
	for _, ele := range a {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	fmt.Println(len(uniq))
}
