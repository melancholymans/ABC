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
	r1 := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r1[0])
	q, _ := strconv.Atoi(r1[1])
	ad := make([]int, n)
	for i := 0; i < q; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		l, _ := strconv.Atoi(r2[0])
		r, _ := strconv.Atoi(r2[1])
		t, _ := strconv.Atoi(r2[2])
		for i := l - 1; i < r; i++ {
			ad[i] = t
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, ad[i])
	}
}
