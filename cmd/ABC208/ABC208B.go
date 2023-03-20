package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	p, _ := strconv.Atoi(s)
	count := 0
	for i := 10; i > 0; i-- {
		f := factorial(i)
		for {
			if p-f < 0 {
				break
			} else {
				p = p - f
				count += 1
			}
		}
	}
	fmt.Fprintln(writer, count)
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}
