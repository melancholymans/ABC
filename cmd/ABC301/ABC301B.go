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
	sl := nis(n)
	fl := make([]int, 0)
	for i := 0; i < n; i++ {
		if i+1 < n {
			fl = append(fl, makeSequence(sl[i], sl[i+1])...)
		} else {
			fl = append(fl, sl[i])
		}
	}
	for i := 0; i < len(fl); i++ {
		fmt.Fprint(wtr, fl[i], " ")
	}
	fmt.Fprintln(wtr, " ")
}

func makeSequence(a, b int) []int {
	ms := make([]int, Abs(b-a))
	if b-a > 0 {
		for i := 0; i < (b - a); i++ {
			ms[i] = a + i
		}
	} else {
		for i := 0; i < a-b; i++ {
			ms[i] = a - i
		}
	}
	return ms
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
