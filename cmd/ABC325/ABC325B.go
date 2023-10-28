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

type IntPart [2]int

func main() {
	defer wtr.Flush()
	n := ni()
	sl := make([]IntPart, n)
	for i := 0; i < n; i++ {
		sl[i][0], sl[i][1] = ni2() //0が人数、1が時差
	}
	mp := make([]int, 24)
	for i := 0; i < n; i++ {
		for j := 9; j < 18; j++ {
			mp[(j+sl[i][1])%24] += sl[i][0]
		}
	}
	rst := 0
	for _, v := range mp {
		rst = max(rst, v)
	}
	fmt.Fprintln(wtr, rst)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
