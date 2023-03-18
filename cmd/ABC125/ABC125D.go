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
	p := make([]byte, 10000000)
	sc.Buffer(p, 10000000)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	a := strings.Split(sc.Text(), " ")
	ad := make([]int, n)
	minus := 0
	var total int
	min := 1000000000
	for i := 0; i < n; i++ {
		ad[i], _ = strconv.Atoi(a[i])
		if ad[i] < 0 {
			minus += 1
		}
		m := Abs(ad[i])
		total += m
		if m < min {
			min = m
		}
	}
	if minus%2 != 0 {
		fmt.Fprintln(writer, total-2*min)
	} else {
		fmt.Fprintln(writer, total)
	}
}

func Abs(a int) int {
	if a >= 0 {
		return +a
	}
	return -a
}
