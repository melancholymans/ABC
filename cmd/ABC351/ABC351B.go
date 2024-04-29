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

func ns() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wtr.Flush()
	n := ni()
	al := make([]string, n*n)
	for i := 0; i < n; i++ {
		s := ns()
		for j := 0; j < n; j++ {
			al[i*n+j] = string(s[j])
		}
	}
	bl := make([]string, n*n)
	for i := 0; i < n; i++ {
		s := ns()
		for j := 0; j < n; j++ {
			bl[i*n+j] = string(s[j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if al[i*n+j] != bl[i*n+j] {
				fmt.Fprintln(wtr, i+1, j+1)
				return
			}
		}
	}
}
