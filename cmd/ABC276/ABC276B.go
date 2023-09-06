package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

//type IntPair [2]int

func main() {
	defer wtr.Flush()
	n, m := ni2()
	sl := make(map[int][]int, 0)
	for i := 0; i < m; i++ {
		k := ni()
		v := ni()
		sl[k] = append(sl[k], v)
		sl[v] = append(sl[v], k)
	}
	for i := 1; i <= n; i++ {
		sort.Ints(sl[i])
	}
	for i := 1; i <= n; i++ {
		if v, ok := sl[i]; ok {
			fmt.Fprint(wtr, len(sl[i]), " ")
			for j := 0; j < len(v); j++ {
				fmt.Fprint(wtr, v[j], " ")
			}
			fmt.Fprintln(wtr, " ")
		} else {
			fmt.Fprintln(wtr, 0)
		}
	}
}
