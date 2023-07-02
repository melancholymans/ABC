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

type IntPair [2]int

func main() {
	defer wtr.Flush()
	n := ni()
	sl := make([]IntPair, 0)
	for i := 0; i < n; i++ {
		var d IntPair
		d[0], d[1] = ni2()
		sl = append(sl, d)
	}
	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			z := float64(sl[j][1]-sl[i][1]) / float64(sl[j][0]-sl[i][0])
			if z >= -1.0 && z <= 1.0 {
				count += 1
			}
		}
	}
	fmt.Fprintln(wtr, count)
}
