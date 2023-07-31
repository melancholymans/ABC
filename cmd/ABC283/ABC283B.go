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
	n := ni()
	a := nis(n)
	q := ni()
	sl := make([][]int, q)
	for i := 0; i < q; i++ {
		if ni() == 1 {
			sl[i] = make([]int, 3)
			sl[i][0] = 1
			sl[i][1] = ni()
			sl[i][2] = ni()
		} else {
			sl[i] = make([]int, 2)
			sl[i][0] = 2
			sl[i][1] = ni()
		}
	}
	for i := 0; i < q; i++ {
		if sl[i][0] == 1 {
			a[sl[i][1]-1] = sl[i][2]
		} else {
			fmt.Fprintln(wtr, a[sl[i][1]-1])
		}
	}
}
