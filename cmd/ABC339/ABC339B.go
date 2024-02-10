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

type IntPar [2]int

func main() {
	defer wtr.Flush()
	h, w, n := ni3()
	mp := make([][]int, h)
	for i := 0; i < h; i++ {
		mp[i] = make([]int, w)
	}
	dx := []int{0, 1, 0, -1} //up,right,down,left 時計方向
	dy := []int{-1, 0, 1, 0}
	tx := 0
	ty := 0
	td := 0
	for i := 0; i < n; i++ {
		if mp[ty][tx] == 0 {
			mp[ty][tx] = 1
			td += 1
		} else {
			mp[ty][tx] = 0
			td -= 1
		}
		td += 4
		td %= 4
		tx += dx[td] + w
		ty += dy[td] + h
		tx %= w
		ty %= h
	}
	for _, v := range mp {
		// 文字列の連結に s+=tのようにすると時間がかかるのであらかじめstringの配列を作っておいてstrings.Joinメソッドで連結する
		r := make([]string, len(v))
		for j, v2 := range v {
			if v2 == 0 {
				r[j] = "."
			} else {
				r[j] = "#"
			}
		}
		fmt.Fprintln(wtr, strings.Join(r, ""))
	}
}
