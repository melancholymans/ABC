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
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				sl[i][j] = 1
				fmt.Println("j==0", i, j, sl[i][j])
				continue
			}
			if i == j {
				sl[i][j] = 1
				fmt.Println("i==j", i, j, sl[i][j])
				continue
			}
			sl[i][j] = sl[i-1][j-1] + sl[i-1][j]
			fmt.Println(i, j, sl[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Fprint(wtr, sl[i][j], " ")
		}
		fmt.Fprintln(wtr, " ")
	}
}
