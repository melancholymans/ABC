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
	m, _ := strconv.Atoi(s[1])
	ad := make(map[int][]int, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		l := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(l[0])
		b, _ := strconv.Atoi(l[1])
		ad[a] = append(ad[a], b)
		ad[b] = append(ad[b], a)
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, len(ad[i]))
	}
}
