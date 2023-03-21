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
	k, _ := strconv.Atoi(s[2])
	l := b - a + 1
	if l < k {
		k = l
	}
	if l-2*k < 0 {
		k = k + (l - 2*k)
	}
	for i := a; i < a+k; i++ {
		fmt.Fprintln(writer, i)
	}
	k, _ = strconv.Atoi(s[2])
	if l < k {
		k = l
	}
	for i := l - k + a; i < a+l; i++ {
		fmt.Fprintln(writer, i)
	}
}
