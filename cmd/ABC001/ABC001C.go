package main

import (
	"bufio"
	"fmt"
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
		b, e := os.ReadFile("./input")
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

type IntPart [2]int

func main() {
	defer wtr.Flush()
	n := ni()
	p := wpow(ni())
	fmt.Fprintln(wtr, direct(n, p), p)
}

func direct(a, b int) string {
	if b == 0 {
		return "C"
	}
	switch a = a * 10; {
	case 1125 <= a && 3375 > a:
		return "NNE"
	case 3375 <= a && 5625 > a:
		return "NE"
	case 5625 <= a && 7875 > a:
		return "ENE"
	case 7875 <= a && 10125 > a:
		return "E"
	case 10125 <= a && 12375 > a:
		return "ESE"
	case 12375 <= a && 14625 > a:
		return "SE"
	case 14625 <= a && 16875 > a:
		return "SSE"
	case 16875 <= a && 19125 > a:
		return "S"
	case 19125 <= a && 21375 > a:
		return "SSW"
	case 21375 <= a && 23625 > a:
		return "SW"
	case 23625 <= a && 25875 > a:
		return "WSW"
	case 25875 <= a && 28125 > a:
		return "W"
	case 28125 <= a && 30375 > a:
		return "WNW"
	case 30375 <= a && 32625 > a:
		return "NW"
	case 32625 <= a && 34875 > a:
		return "NNW"
	default:
		return "N"
	}
}

func wpow(a int) int {
	f := float64(a) / 60.0
	f = f * 10
	switch w := int(math.Round(f)); {
	case 0 <= w && 2 >= w:
		return 0
	case 3 <= w && 15 >= w:
		return 1
	case 16 <= w && 33 >= w:
		return 2
	case 34 <= w && 54 >= w:
		return 3
	case 55 <= w && 79 >= w:
		return 4
	case 80 <= w && 107 >= w:
		return 5
	case 108 <= w && 138 >= w:
		return 6
	case 139 <= w && 171 >= w:
		return 7
	case 172 <= w && 207 >= w:
		return 8
	case 208 <= w && 244 >= w:
		return 9
	case 245 <= w && 284 >= w:
		return 10
	case 285 <= w && 326 >= w:
		return 11
	default:
		return 12
	}
}
