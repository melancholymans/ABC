package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	var total int
	fmt.Scan(&n, &a, &b)
	for i := 1; i <= n; i++ {
		r := sub(i, a, b)
		total += r
	}
	fmt.Println(total)
}

func sub(n, a, b int) int {
	s := strconv.Itoa(n)
	l := len(s)
	var sum int
	for i := 0; i < l; i++ {
		sum += int(s[i]) - 48
	}
	if (a <= sum) && (sum <= b) {
		return n
	}
	return 0
}
