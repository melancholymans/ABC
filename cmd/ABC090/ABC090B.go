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
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	result := make([]int, 0)
	for i := a; i <= b; i++ {
		n, _ := strconv.Atoi(reverse(strconv.Itoa(i)))
		if i == n {
			result = append(result, i)
		}
	}
	fmt.Fprintln(writer, len(result))
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
