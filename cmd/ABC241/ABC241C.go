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

type point struct {
	r int
	w int
}

func main() {
	defer wtr.Flush()
	n := ni()
	m := make([][]bool, n)
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
		s := ns()
		for j, v := range s {
			if string(v) == "#" {
				m[i][j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p := point{i, j}
			dr := [4][]point{}
			dr[0] = []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}
			dr[1] = []point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}}
			dr[2] = []point{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
			dr[3] = []point{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}}
			for k := 0; k < 4; k++ {
				count := 0
				for _, d := range dr[k] {
					np := pointAdd(p, d)
					if !np.isValid(n, n) {
						count = 0
						break
					}
					if m[np.r][np.w] {
						count += 1
					}
				}
				if count >= 4 {
					fmt.Fprintln(wtr, "Yes")
					return
				}
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}

func (p point) isValid(r, w int) bool {
	return (0 <= p.r && p.r < r) && (0 <= p.w && p.w < w)
}

func pointAdd(a, b point) point {
	return point{a.r + b.r, a.w + b.w}
}
