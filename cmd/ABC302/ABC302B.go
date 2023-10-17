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

type point struct {
	x int
	y int
}

func (p point) isvalid(x, y int) bool {
	return (0 <= p.x && p.x < x) && (0 <= p.y && p.y < y)
}

func main() {
	defer wtr.Flush()
	h, w := ni2()
	sl := make([][]string, h)
	for i := 0; i < h; i++ {
		sl[i] = strings.Split(ns(), "")
	}
	dx := [8]int{1, 1, 0, -1, -1, -1, 0, 1}
	dy := [8]int{0, 1, 1, 1, 0, -1, -1, -1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 8; k++ {
				t := ""
				for l := 0; l < 5; l++ {
					pt := point{i + dx[k]*l, j + dy[k]*l}
					if pt.isvalid(h, w) {
						t += sl[pt.x][pt.y]
					}
				}
				if t == "snuke" {
					for l := 0; l < 5; l++ {
						pt := point{i + dx[k]*l, j + dy[k]*l}
						fmt.Fprintf(wtr, "%d %d\n", pt.x+1, pt.y+1)
					}
					return
				}
			}
		}
	}
}
