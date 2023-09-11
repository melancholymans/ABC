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
	if n < 1000 {
		fmt.Fprintln(wtr, n)
	} else if 1000 <= n && n < 10000 {
		fmt.Fprintln(wtr, n/10*10)
	} else if 10000 <= n && n < 100000 {
		fmt.Fprintln(wtr, n/100*100)
	} else if 100000 <= n && n < 1000000 {
		fmt.Fprintln(wtr, n/1000*1000)
	} else if 1000000 <= n && n < 10000000 {
		fmt.Fprintln(wtr, n/10000*10000)
	} else if 10000000 <= n && n < 100000000 {
		fmt.Fprintln(wtr, n/100000*100000)
	} else if 100000000 <= n && n < 1000000000 {
		fmt.Fprintln(wtr, n/1000000*1000000)
	}
}
