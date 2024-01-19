package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func ns() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wtr.Flush()
	n, k := ni2()
	sl := nis(n)
	sp := make([]int, n)
	for i := 0; i < n; i++ {
		sp[i] = sl[i]
	}
	kd := k % n
	kn := k / n
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	th := -1
	if kd != 0 {
		th = sl[kd-1]
	}
	for _, v := range sp {
		if v <= th {
			fmt.Fprintln(wtr, kn+1)
		} else {
			fmt.Fprintln(wtr, kn)
		}
	}
}
