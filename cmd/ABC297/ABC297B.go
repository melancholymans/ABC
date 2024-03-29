package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
		b, e := ioutil.ReadFile("./input")
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
	s := ns()
	ki := 0
	r1i, r2i := 0, 0
	b1i, b2i := 0, 0
	for i := 0; i < 8; i++ {
		switch string(s[i]) {
		case "K":
			ki = i + 1
		case "R":
			if r1i == 0 {
				r1i = i + 1
			} else {
				r2i = i + 1
			}
		case "B":
			if b1i == 0 {
				b1i = i + 1
			} else {
				b2i = i + 1
			}
		}
	}
	if ki < r1i || ki > r2i {
		fmt.Fprintln(wtr, "No")
		return
	}
	if b1i%2 == b2i%2 {
		fmt.Fprintln(wtr, "No")
		return
	}
	fmt.Fprintln(wtr, "Yes")
}
