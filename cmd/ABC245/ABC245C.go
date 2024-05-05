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
	n, k := ni2()
	as := nis(n)
	bs := nis(n)
	ta := true
	tb := true
	for i := 1; i < n; i++ {
		var nta, ntb bool
		if ta {
			if abs(as[i-1]-as[i]) <= k {
				nta = true
			}
			if abs(as[i-1]-bs[i]) <= k {
				ntb = true
			}
		}
		if tb {
			if abs(bs[i-1]-as[i]) <= k {
				nta = true
			}
			if abs(bs[i-1]-bs[i]) <= k {
				ntb = true
			}
		}
		ta = nta
		tb = ntb
	}
	if ta || tb {
		fmt.Fprintln(wtr, "Yes")
	} else {
		fmt.Fprintln(wtr, "No")

	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
