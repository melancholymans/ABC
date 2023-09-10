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
	n := ni()
	sa := make([][]int, n)
	for i := 0; i < n; i++ {
		sa[i] = nis(n)
	}
	sb := make([][]int, n)
	for i := 0; i < n; i++ {
		sb[i] = nis(n)
	}
	//回転なし
	if cmp(n, sb, sa) {
		fmt.Fprintln(wtr, "Yes")
		return
	}
	if cmp90(n, sb, sa) {
		fmt.Fprintln(wtr, "Yes")
		return
	}
	if cmp180(n, sb, sa) {
		fmt.Fprintln(wtr, "Yes")
		return
	}
	if cmp270(n, sb, sa) {
		fmt.Fprintln(wtr, "Yes")
		return
	}
	fmt.Fprintln(wtr, "No")
}

func cmp(n int, a, b [][]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[i][j] == 1 {
				if a[i][j] != b[i][j] {
					return false
				}
			}
		}
	}
	return true
}

func cmp90(n int, a, b [][]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[n-1-j][i] == 1 {
				if a[i][j] != b[n-1-j][i] {
					return false
				}
			}
		}
	}
	return true
}

func cmp180(n int, a, b [][]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[n-1-i][n-1-j] == 1 {
				if a[i][j] != b[n-1-i][n-1-j] {
					return false
				}
			}
		}
	}
	return true
}

func cmp270(n int, a, b [][]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[j][n-1-i] == 1 {
				if a[i][j] != b[j][n-1-i] {
					return false
				}
			}
		}
	}
	return true
}
