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

type IntPair [2]int

func main() {
	defer wtr.Flush()
	s := []byte(ns())
	n := len(s)
	mp := map[byte]int{}
	for i := 0; i < n; i++ {
		mp[s[i]] += 1
	}
	sl := make([]IntPair, n)
	//[0] - key [1] - cont
	i := 0
	for k, v := range mp {
		sl[i][0] = int(k)
		sl[i][1] = v
		i += 1
	}
	sort.Slice(sl, func(i, j int) bool {
		if sl[i][1] == sl[j][1] {
			return sl[i][0] < sl[j][0]
		}
		return sl[i][1] > sl[j][1]
	})
	fmt.Fprintln(wtr, string(sl[0][0]))
}
