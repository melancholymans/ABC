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

func main() {
	defer wtr.Flush()
	n, m := ni2()
	edge1 := make([][]bool, n)
	edge2 := make([][]bool, n)
	for i := 0; i < n; i++ {
		edge1[i] = make([]bool, n)
		edge2[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		s := ni() - 1
		t := ni() - 1
		edge1[s][t] = true
		edge1[t][s] = true
	}
	for i := 0; i < m; i++ {
		s := ni() - 1
		t := ni() - 1
		edge2[s][t] = true
		edge2[t][s] = true
	}
	ls := make([]int, n)
	for i := 0; i < n; i++ {
		ls[i] = i
	}
	for {
		r := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if edge1[i][j] {
					if edge2[ls[i]][ls[j]] {
					} else {
						r = false
					}
				}
			}
		}
		if r {
			fmt.Fprintln(wtr, "Yes")
			return
		}
		if !nextPermutation(sort.IntSlice(ls)) {
			break
		}
	}
	fmt.Fprintln(wtr, "No")
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
