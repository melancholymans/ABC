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
	n, m := ni2()
	sl := nis(m)
	mp := make(map[int]bool)
	for _, v := range sl {
		mp[v] = true
	}
	st := 0
	en := 0
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i + 1
	}
	for i := 1; i <= n; i++ {
		if st == 0 && mp[i] {
			st = i
			en = i + 1
			continue
		} else if mp[i] {
			if en == i {
				en++
				continue
			}
		}
		if st == 0 {
			continue
		}
		t := en
		for j := st; j <= en; j++ {
			r[j-1] = t
			t--
		}
		st = 0
		en = 0
	}
	sj := make([]string, 0)
	for i := 0; i < len(r); i++ {
		sj = append(sj, strconv.Itoa(r[i]))
	}
	fmt.Fprintln(wtr, strings.Join(sj, " "))
}
