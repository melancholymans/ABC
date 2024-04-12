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

const (
	asc sortOrder = iota
	desc
)

type point struct {
	x int
	y int
}

type sortOrder int

type Sort2ArOptions struct {
	keys   []int
	orders []sortOrder
}

type Sort2ArOption func(*Sort2ArOptions)

func main() {
	defer wtr.Flush()
	n := ni()
	xs := make([]int, n)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i], ys[i] = ni2()
	}
	rs := make([][][2]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			t := pointDistDouble(point{xs[i], ys[i]}, point{xs[j], ys[j]})
			rs[i] = append(rs[i], [2]int{t, j})
			rs[j] = append(rs[j], [2]int{t, i})
		}
	}
	for i := 0; i < n; i++ {
		sort2ar(rs[i], opt2ar(0, desc), opt2ar(1, asc))
		fmt.Fprintln(wtr, rs[i][0][1]+1)
	}
}

func pointDistDouble(a, b point) int {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
}

func sort2ar(sl [][2]int, setters ...Sort2ArOption) {
	args := &Sort2ArOptions{}
	for _, setter := range setters {
		setter(args)
	}
	sort.Slice(sl, func(i, j int) bool {
		for idx, key := range args.keys {
			if sl[i][key] == sl[j][key] {
				continue
			}
			switch args.orders[idx] {
			case asc:
				return sl[i][key] < sl[j][key]
			case desc:
				return sl[i][key] > sl[j][key]
			}
		}
		return true
	})
}

func opt2ar(key int, order sortOrder) Sort2ArOption {
	return func(args *Sort2ArOptions) {
		args.keys = append(args.keys, key)
		args.orders = append(args.orders, order)
	}
}
