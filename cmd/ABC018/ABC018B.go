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
	s := sc.Text()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		rev := strings.Split(sc.Text(), " ")
		l, _ := strconv.Atoi(rev[0])
		r, _ := strconv.Atoi(rev[1])
		s = s[0:l-1] + Reverses(s[l-1:r]) + s[r:]
	}
	fmt.Fprintln(writer, s)
}

func Reverses(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
