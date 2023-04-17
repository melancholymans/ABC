package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	x, _ := strconv.Atoi(r1[1])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	pd := make([]int, 0)
	for i := 0; i < n; i++ {
		p, _ := strconv.Atoi(r2[i])
		pd = append(pd, p)
	}
	total := 0
	ft := "%0" + strconv.Itoa(n) + "s"
	s := fmt.Sprintf(ft, Debugbit(x))
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			total += pd[n-i-1]
		}
	}
	fmt.Fprintln(writer, total)
}

func Debugbit(n int) string {
	r := ""
	for i := Bitlen(n) - 1; i >= 0; i-- {
		if n&(1<<i) != 0 {
			r += "1"
		} else {
			r += "0"
		}
	}
	return r
}

func Bitlen(a int) int {
	return bits.Len(uint(a))
}
