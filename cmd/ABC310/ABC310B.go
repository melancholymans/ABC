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
	n, m := ni2()
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, 1+m)
		sl[i][0] = ni() //price
		c := ni()
		for j := 0; j < c; j++ {
			f := ni()
			sl[i][f] = 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if sl[i][0] < sl[j][0] {
				continue
			}
			flag := true
			total := 0
			for k := 1; k <= m; k++ {
				if sl[j][k] == 1 && sl[i][k] == 0 {
					total += 1
				}
				if sl[i][k] == 1 && sl[j][k] == 0 {
					flag = false
				}
			}
			if !flag {
				continue
			}
			if sl[i][0] > sl[j][0] {
				fmt.Fprintln(wtr, "Yes")
				return
			} else if total > 0 {
				fmt.Fprintln(wtr, "Yes")
				return
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}
