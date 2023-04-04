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
	n, _ := strconv.Atoi(s)
	sl := Divisor(n)
	fmt.Fprintln(writer, len(sl))
	for i := 0; i < len(sl); i++ {
		fmt.Fprintln(writer, sl[i])
	}
}

func Divisor(n int) []int {
	sl := make([]int, 0)
	cd := [4]int{8, 4, 2, 1}
	for _, c := range cd {
		for {
			if n-c < 0 {
				break
			}
			sl = append(sl, c)
			n -= c
		}
	}
	return sl
}
