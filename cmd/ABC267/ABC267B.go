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
	s := strings.Split(ns(), "")
	si := [10]int{}
	for i := 0; i < 10; i++ {
		si[i], _ = strconv.Atoi(s[i])
	}
	p := [7]int{}
	p[0] = si[6]
	p[1] = si[3]
	p[2] = si[1] + si[7]
	p[3] = si[0] + si[4]
	p[4] = si[2] + si[8]
	p[5] = si[5]
	p[6] = si[9]
	if si[0] != 0 {
		fmt.Fprintln(wtr, "No")
		return
	}
	for i := 0; i < 7; i++ {
		if p[i] != 0 {
			if i+1 < 7 && p[i+1] == 0 {
				if findnozero(i+2, p) {
					fmt.Fprintln(wtr, "Yes")
					return
				}
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}

func findnozero(idx int, p [7]int) bool {
	for i := idx; i < 7; i++ {
		if p[i] != 0 {
			return true
		}
	}
	return false
}
