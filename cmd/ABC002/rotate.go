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
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = ni()
		}
	}
	//通常
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprint(wtr, sl[i][j], " ")
		}
		fmt.Fprintln(wtr, " ")
	}
	//時計方向90度回転
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(wtr, sl[n-1-j][i], " ")
		}
		fmt.Fprintln(wtr, " ")
	}
	//時計方向180度回転
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprint(wtr, sl[n-1-i][m-1-j], " ")
		}
		fmt.Fprintln(wtr, " ")
	}
	//時計方向270度回転
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(wtr, sl[j][m-1-i], " ")
		}
		fmt.Fprintln(wtr, " ")
	}
}