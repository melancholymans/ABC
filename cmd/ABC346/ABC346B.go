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
	w, b := ni2()
	s := "wbwbwwbwbwbw"
	sn := len(s)
	for i := 0; i < 100; i++ {
		s += "wbwbwwbwbwbw"
	}
	tn := len(s)
	sl := make([]byte, tn)
	for i := 0; i < tn; i++ {
		sl[i] = s[i]
	}
	for i := 0; i <= sn; i++ {
		var tb, tw int
		for j := i; j < tn; j++ {
			if sl[j] == 98 {
				tb += 1
			} else {
				tw += 1
			}
			if tb == b && tw == w {
				fmt.Fprintln(wtr, "Yes")
				return
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}
