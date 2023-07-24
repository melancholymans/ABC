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

type StringPair [2]string

func main() {
	defer wtr.Flush()
	n := ni()
	sl := make([]StringPair, n)
	mm := map[string]int{}
	for i := 0; i < n; i++ {
		sl[i][0] = ns()
		sl[i][1] = ns()
		mm[sl[i][0]] += 1
		mm[sl[i][1]] += 1
	}
	for i := 0; i < n; i++ {
		if mm[sl[i][0]] == 1 || mm[sl[i][1]] == 1 {
		} else {
			if sl[i][0] == sl[i][1] && mm[sl[i][1]] == 2 {
			} else {
				fmt.Fprintln(wtr, "No")
				return
			}
		}
	}
	fmt.Fprintln(wtr, "Yes")
}
