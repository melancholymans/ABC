package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r1 := sc.Text()
	n, _ := strconv.Atoi(r1)
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	sl := make([]int, 0)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		sl = append(sl, a)
	}
	count := 0
	for i := 0; i < len(sl)-2; i++ {
		for j := i + 1; j < len(sl)-1; j++ {
			for k := j + 1; k < len(sl); k++ {
				if sl[i] == sl[j] || sl[i] == sl[k] || sl[j] == sl[k] {
					continue
				}
				if calc(sl[i], sl[j], sl[k]) {
					count += 1
				}
			}
		}
	}
	fmt.Fprintln(writer, count)
}

func calc(a, b, c int) bool {
	m := Max(Max(a, b), c)
	l := a + b + c
	if (l - m) > m {
		return true
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
