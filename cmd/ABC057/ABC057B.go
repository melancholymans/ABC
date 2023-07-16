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

type IntPair [2]int

func main() {
	defer wtr.Flush()
	n, m := ni2()
	sg := make([]IntPair, n)
	sc := make([]IntPair, m)
	for i := 0; i < n; i++ {
		sg[i][0], sg[i][1] = ni2()
	}
	for i := 0; i < m; i++ {
		sc[i][0], sc[i][1] = ni2()
	}
	idx := 0
	for i := 0; i < n; i++ {
		mmin := math.MaxInt64
		for j := 0; j < m; j++ {
			d := calc(sg[i][0], sg[i][1], sc[j][0], sc[j][1])
			if d < mmin {
				mmin = d
				idx = j + 1
			}
		}
		fmt.Fprintln(wtr, idx)
	}
}

func calc(a, b, c, d int) int {
	return Abs(a-c) + Abs(b-d)
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
