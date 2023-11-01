package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := os.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}
func ni2() (int, int) {
	return ni(), ni()
}
func ni3() (int, int, int) {
	return ni(), ni(), ni()
}
func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}
func ns() string {
	sc.Scan()
	return sc.Text()
}
func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func main() {
	defer wtr.Flush()
	t := ni()
	n := ni()
	sa := nis(n)
	m := ni()
	sb := nis(m)
	if n < m {
		fmt.Fprintln(wtr, "no")
		return
	}
	flag := true
	idx := 0
	for i := 0; i < m; i++ {
		if idx == n {
			flag = false
			break
		}
		for {
			if idx == n {
				flag = false
				break
			}
			if sa[idx] <= sb[i] && sb[i] <= sa[idx]+t {
				idx += 1
				break
			} else {
				idx += 1
			}
		}
	}
	if flag {
		fmt.Fprintln(wtr, "yes")
	} else {
		fmt.Fprintln(wtr, "no")
	}
}
