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
	n, k := ni2()
	t := []byte(ns())
	sl := make([]int, n)
	for i := 0; i < len(t); i++ {
		sl[i] = int(t[i] - 97)
	}
	bs := make([]int, 26)
	tm := make([]int, 26)
	for i := 0; i < n; i++ {
		bs[sl[i]] += 1
		tm[sl[i]] += 1
	}
	ts := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			if tm[j] == 0 {
				continue
			}
			if sl[i] == j {
				ts[i] = j
				break
			}
			bs[sl[i]] -= 1
			tm[j] -= 1
			df := 0
			for k := 0; k < 26; k++ {
				if bs[k] > tm[k] {
					df += bs[k] - tm[k]
				}
			}
			bs[sl[i]] += 1
			tm[j] += 1
			if df <= k-1 {
				k -= 1
				ts[i] = j
				break
			}
		}
		bs[sl[i]] -= 1
		tm[ts[i]] -= 1
	}
	rst := ""
	for i := 0; i < len(ts); i++ {
		rst += string(ts[i] + 97)
	}
	fmt.Fprintln(wtr, rst)
}
