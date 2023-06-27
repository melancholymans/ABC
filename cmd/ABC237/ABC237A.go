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
	n := ns()
	sig := 1
	if string(n[0]) == "-" {
		sig = -1
		n = n[1:]
	}
	var ln uint64
	ln, _ = strconv.ParseUint(n, 10, 64)
	if sig == 1 {
		if ln >= Pow(2, 31) {
			fmt.Fprintln(wtr, "No")
		} else {
			fmt.Fprintln(wtr, "Yes")
		}
	} else if sig == -1 {
		if ln > Pow(2, 31) {
			fmt.Fprintln(wtr, "No")
		} else {
			fmt.Fprintln(wtr, "Yes")
		}
	}
}

func Pow(a, b int) uint64 {
	return uint64(math.Pow(float64(a), float64(b)))
}
