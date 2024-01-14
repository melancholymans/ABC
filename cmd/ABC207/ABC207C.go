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
	sl := make([][2]int, n)
	count := 0
	for i := 0; i < n; i++ {
		t, l, r := ni3()
		l = l * 3
		r = r * 3
		switch t {
		case 1:
		case 2:
			r -= 1
		case 3:
			l += 1
		case 4:
			l += 1
			r -= 1
		}
		sl[i][0] = l
		sl[i][1] = r
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if sl[i][1] >= sl[j][0] && sl[i][0] <= sl[j][1] {
				count += 1
			} else if sl[j][1] >= sl[i][0] && sl[j][0] <= sl[i][1] {
				count += 1
			}
		}
	}
	fmt.Fprintln(wtr, count)
}
