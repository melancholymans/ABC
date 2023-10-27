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

type IntPair [2]int

func main() {
	defer wtr.Flush()
	n := ni()
	sl := make([][]byte, n)
	lank := make([]IntPair, n)
	for i := 0; i < n; i++ {
		sl[i] = []byte(ns())
		for j := 0; j < n; j++ {
			if sl[i][j] == byte(111) {
				lank[i][0] += 1
			}
		}
		lank[i][1] = i + 1
	}
	sort.Slice(lank, func(i, j int) bool {
		if lank[i][0] == lank[j][0] { //テストの点数が一緒なら
			return lank[i][1] < lank[j][1] //playerの番号が小さいほうか選択
		}
		return lank[i][0] > lank[j][0] //テストの点数が大きいほう
	})
	for i := 0; i < n; i++ {
		fmt.Fprint(wtr, lank[i][1], " ")
	}
	fmt.Fprintln(wtr, " ")
}
