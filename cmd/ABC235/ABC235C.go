package main

import (
	"bufio"
	"fmt"
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
	n, q := ni2()
	al := nis(n)
	mp := make(map[int][]int)
	for i := 0; i < n; i++ {
		if _, ok := mp[al[i]]; !ok {
			mp[al[i]] = []int{i + 1}
		} else {
			mp[al[i]] = append(mp[al[i]], i+1)
		}
	}
	for i := 0; i < q; i++ {
		x := ni()
		k := ni() - 1
		if v, ok := mp[x]; ok {
			if k < len(v) {
				fmt.Fprintln(wtr, v[k])
			} else {
				fmt.Fprintln(wtr, -1)
			}
		} else {
			fmt.Fprintln(wtr, -1)
		}
	}
}
