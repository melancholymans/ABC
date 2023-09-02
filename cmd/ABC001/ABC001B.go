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
	m := ni()
	if m < 0.1*1000 {
		fmt.Fprintf(wtr, "%02d\n", 0)
	} else if 0.1*1000 <= m && m <= 5*1000 {
		fmt.Fprintf(wtr, "%02d\n", m/100)
	} else if 6*1000 <= m && m <= 30*1000 {
		fmt.Fprintf(wtr, "%02d\n", (m+50*1000)/1000)
	} else if 35*1000 <= m && m <= 70*1000 {
		fmt.Fprintf(wtr, "%02d\n", ((m-30*1000)/5+80*1000)/1000)
	} else if 70*1000 < m {
		fmt.Fprintf(wtr, "%02d\n", 89)
	}
}
