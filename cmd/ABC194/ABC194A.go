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
func ns() string {
	sc.Scan()
	return sc.Text()
}
func main() {
	defer wtr.Flush()
	a, b := ni2()
	k := a + b
	s := b
	if k >= 15 && s >= 8 {
		fmt.Fprintln(wtr, 1)
	} else if k >= 10 && s >= 3 {
		fmt.Fprintln(wtr, 2)
	} else if k >= 3 {
		fmt.Fprintln(wtr, 3)
	} else {
		fmt.Fprintln(wtr, 4)
	}
}
