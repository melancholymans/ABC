package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type st struct {
	id   int
	win  int
	hand string
}

func main() {
	defer wtr.Flush()
	n, m := ni2()
	ss := make([]st, 2*n)
	for i := 0; i < 2*n; i++ {
		s := ns()
		ss[i] = st{i, 0, s}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f := string(ss[2*j].hand[i])
			t := string(ss[2*j+1].hand[i])
			if t == f {
			} else if (f == "G" && t == "C") || (f == "C" && t == "P") || (f == "P" && t == "G") {
				ss[2*j].win++
			} else {
				ss[2*j+1].win++
			}
		}
		sort.Slice(ss, func(i, j int) bool {
			if ss[i].win == ss[j].win {
				return ss[i].id < ss[j].id
			}
			return ss[i].win > ss[j].win
		})
	}
	for i := 0; i < 2*n; i++ {
		fmt.Fprintln(wtr, ss[i].id+1)
	}
}
