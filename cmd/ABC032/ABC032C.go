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
	mmax := 0
	n, k := ni2()
	ns := make([]int, n)
	flag := false
	for i := 0; i < n; i++ {
		ns[i] = ni()
		if ns[i] == 0 {
			flag = true
		}
	}
	if flag {
		fmt.Fprintln(wtr, n)
		return
	}
	if k == 0 {
		fmt.Fprintln(wtr, 0)
		return
	}
	s := 0
	t := 0
	sum := 1
	for {
		for t != n {
			if sum*ns[t] > k {
				break
			}
			sum *= ns[t]
			t++
		}
		mmax = max(mmax, t-s)
		if t == n {
			break
		}
		sum *= ns[t]
		t++
		for s != n {
			if sum <= k {
				break
			}
			sum /= ns[s]
			s++
		}
	}
	fmt.Fprintln(wtr, mmax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
