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
	total := math.MinInt64
	n := ni()
	ns := nis(n)
	for i := 0; i < n; i++ {
		mc := math.MinInt64
		mi := 0
		t := 0
		for j := 0; j < i; j++ {
			t = 0
			for k := j + 1; k <= i; k += 2 {
				t += ns[k]
			}
			if t > mc {
				mc = t
				mi = j
			}
		}
		t = 0
		for j := i + 1; j < n; j++ {
			if (j-i)%2 == 0 {
				continue
			}
			t += ns[j]
			if t > mc {
				mc = t
				mi = j
			}
		}
		t2 := 0
		if mi < i {
			for j := mi; j <= i; j++ {
				if (j-mi)%2 == 0 {
					t2 += ns[j]
				}
			}
		} else {
			for j := i; j <= mi; j++ {
				if (j-i)%2 == 0 {
					t2 += ns[j]
				}
			}
		}
		total = max(total, t2)
	}
	fmt.Fprintln(wtr, total)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
