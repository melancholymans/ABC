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
	x := ns()
	normal := "abcdefghijklmnopqrstuvwxyz"
	conv := make(map[string]int)
	for i, v := range x {
		conv[string(v)] = i
	}
	n := ni()
	ss := make([][2]string, n)
	for i := 0; i < n; i++ {
		s := ns()
		cs := ""
		for _, v := range s {
			idx := conv[string(v)]
			cs += string(normal[idx])
		}
		ss[i][0] = cs
		ss[i][1] = s
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i][0] < ss[j][0]
	})
	for _, v := range ss {
		fmt.Fprintln(wtr, v[1])
	}
}
