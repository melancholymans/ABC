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
	s := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(s[0])
	x, _ := strconv.Atoi(s[1])
	la := make([]int, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())
		la = append(la, m)
	}
	m := make(map[int]bool)
	uniq := []int{}
	for _, ele := range la {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	total := 0
	min := 1000000
	count := 0
	for _, a := range uniq {
		count += 1
		total += a
		if a < min {
			min = a
		}
	}
	fmt.Fprintln(writer, count+(x-total)/min)
}
