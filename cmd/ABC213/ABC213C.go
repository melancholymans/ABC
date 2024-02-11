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

type IntPar [2]int

func main() {
	defer wtr.Flush()
	_, _, n := ni3()
	ns := make([]IntPar, n)
	ons := make([]IntPar, n)
	xm := make(map[int]int)
	ym := make(map[int]int)
	for i := 0; i < n; i++ {
		a, b := ni2()
		ns[i][0], ns[i][1] = a, b
		ons[i][0], ons[i][1] = a, b
	}
	//ns配列を行インデックスを昇順にソートしている
	sort.Slice(ns, func(i, j int) bool {
		return ns[i][0] < ns[j][0]
	})
	//多分cはカラムのc
	c := 0
	t := 0
	for i := 0; i < n; i++ {
		if ns[i][0] != t {
			t = ns[i][0]
			c += 1
		}
		xm[t] = c
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i][1] < ns[j][1]
	})
	c = 0
	t = 0
	for i := 0; i < n; i++ {
		if ns[i][1] != t {
			t = ns[i][1]
			c += 1
		}
		ym[t] = c
	}
	for _, v := range ons {
		fmt.Fprintf(wtr, "%v %v\n", xm[v[0]], ym[v[1]])
	}
}
