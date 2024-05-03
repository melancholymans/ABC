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

type point struct {
	x, y int
}

func main() {
	defer wtr.Flush()
	x1, y1 := ni2()
	x2, y2 := ni2()
	d := [8]point{
		{2, 1},
		{1, 2},
		{-2, 1},
		{-1, 2},
		{-2, -1},
		{-1, -2},
		{1, -2},
		{2, -1},
	}
	s1 := make([]point, 8)
	for i := 0; i < 8; i++ {
		s1[i].x = x1 + d[i].x
		s1[i].y = y1 + d[i].y
	}
	s2 := make([]point, 8)
	for i := 0; i < 8; i++ {
		s2[i].x = x2 + d[i].x
		s2[i].y = y2 + d[i].y
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if s1[i] == s2[j] /*cmp(s1[i], s2[j])*/ {
				fmt.Fprintln(wtr, "Yes")
				return
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}

//わざわざx値とy値を分解して比較しているが構造体ごと比較してもよい
/*
func cmp(a1, a2 point) bool {
	if a1.x == a2.x && a1.y == a2.y {
		return true
	}
	return false
}
*/
