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
	s := ns()
	n := len(s)
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		ns[i] = int(s[i] - 65)
	}
	flag := -1
	ts := make([]int, 0)
	for i := 0; i < n; i++ {
		if flag != ns[i] {
			ts = append(ts, ns[i])
			flag = ns[i]
		}
	}
	if len(ts) > 3 {
		fmt.Fprintln(wtr, "No")
		return
	}
	for i := 0; i < len(ts)-1; i++ {
		if ts[i] > ts[i+1] {
			fmt.Fprintln(wtr, "No")
			return
		}
	}
	fmt.Fprintln(wtr, "Yes")
}
