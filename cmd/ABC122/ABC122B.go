package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := sc.Text()
	m := make([][]rune, 0)
	s := make([]rune, 0)
	for i, v := range r {
		if calc(v) {
			s = append(s, v)
		} else if !calc(v) {
			m = append(m, s)
			s = make([]rune, 0)
		}
		if i == len(r)-1 {
			m = append(m, s)
		}
	}
	maxm := math.MinInt64
	for _, v := range m {
		if maxm < len(v) {
			maxm = len(v)
		}
	}
	fmt.Fprintln(writer, maxm)
}

func calc(s rune) bool {
	switch s {
	case 'A':
		return true
	case 'C':
		return true
	case 'G':
		return true
	case 'T':
		return true
	default:
		return false
	}
}
