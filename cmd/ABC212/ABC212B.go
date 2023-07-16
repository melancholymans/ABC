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
	x := ns()
	n := strings.Split(x, "")
	if n[0] == n[1] && n[1] == n[2] && n[2] == n[3] {
		fmt.Fprintln(wtr, "Weak")
		return
	}
	if cmp(n[0], n[1]) && cmp(n[1], n[2]) && cmp(n[2], n[3]) {
		fmt.Fprintln(wtr, "Weak")
		return
	}
	fmt.Fprintln(wtr, "Strong")
}

func cmp(a, b string) bool {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	if (x+1)%10 == y {
		return true
	} else {
		return false
	}
}
