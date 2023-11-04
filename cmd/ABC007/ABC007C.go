package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func main() {
	defer wtr.Flush()
	h, w := ni2()
	sh, sw := ni()-1, ni()-1
	gh, gw := ni()-1, ni()-1
	sl := make([][]bool, h)
	dist := make([]int, h*w)
	gr := make([][]int, h*w)
	//# 35,. 46
	for i := 0; i < h; i++ {
		sl[i] = make([]bool, w)
		b := []byte(ns())
		for j := 0; j < w; j++ {
			if b[j] == 46 {
				sl[i][j] = true
			}
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !sl[i][j] {
				continue
			}
			p := i*w + j
			dist[p] = -1
			if i > 0 {
				if sl[i-1][j] {
					gr[p] = append(gr[p], p-w)
				}
			}
			if i < h-1 {
				if sl[i+1][j] {
					gr[p] = append(gr[p], p+w)
				}
			}
			if j > 0 {
				if sl[i][j-1] {
					gr[p] = append(gr[p], p-1)
				}
			}
			if j < w-1 {
				if sl[i][j+1] {
					gr[p] = append(gr[p], p+1)
				}
			}
		}
	}
	s := sh*w + sw
	q := []int{}
	q = append(q, s)
	dist[s] = 0
	for {
		if len(q) == 0 {
			break
		}
		t := q[0]
		q = q[1:]
		for _, v := range gr[t] {
			if dist[v] != -1 {
				continue
			}
			q = append(q, v)
			dist[v] = dist[t] + 1
		}
	}
	fmt.Fprintln(wtr, dist[gh*w+gw])
}
